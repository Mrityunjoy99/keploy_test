package database

import (
	"fmt"
	"time"

	"github.com/sample_project/src/common"
	"gopkg.in/errgo.v2/fmt/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

type GormDialect string

const (
	DialectPostgres GormDialect = "DialectPostgres"
)

func Connect(c *common.Config) (*gorm.DB, error) {
	conn, err := connectToDB(c)
	if err != nil {
		return nil, err
	}

	db = conn

	return db, nil
}

func GetInstance() *gorm.DB {
	return db
}

func SetInstance(g *gorm.DB) {
	db = g
}

func CloseDB() error {
	dbInstance := db
	if dbInstance == nil {
		return errors.New("error while closing db as dbInstance is nil")
	}

	closeDbInstance, serr := dbInstance.DB()
	if serr != nil {
		return errors.New("error while closing db as dbInstance.DB() call failed")
	}

	closeDBerr := closeDbInstance.Close()
	if closeDBerr != nil {
		return errors.New("error while closing db as closeDbInstance.Close() call failed")
	}

	return nil
}

type DBConfig struct {
	AppName  string `validate:"required"`
	Env      string `validate:"required"`
	Host     string `validate:"required"`
	Port     int    `validate:"required"`
	DBName   string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`

	MaxIdleConnCount int `validate:"gte=0"`
	MaxOpenConnCount int `validate:"gtefield=MaxIdleConnCount"`
	// ConnMaxIdleTimeSec sets the maximum amount of time in seconds a connection may be reused.
	ConnMaxIdleTimeSec int `validate:"gte=0"`

	// ConnMaxLifeTimeSec sets the maximum amount of time in seconds a connection may be reused
	ConnMaxLifeTimeSec int `validate:"gtefield=ConnMaxIdleTimeSec"`

	Dialect GormDialect `validate:"required"`

	// GORM perform single create, update, delete operations in transactions by default to ensure database data integrity
	// You can disable it by setting `SkipDefaultTransaction` to true
	SkipDefaultTransaction bool `default:"true"`

	Options *gorm.Config
}

func connectToDB(c *common.Config) (*gorm.DB, error) {
	const (
		RepositoryPostgresConnectWaitSec = 3
		ConnectRetryAttempts             = 10
	)

	dbConfig := DBConfig{
		AppName:            c.App.Name,
		Env:                "local",
		Host:               c.DB.Host,
		Port:               c.DB.Port,
		DBName:             c.DB.Name,
		User:               c.DB.User,
		Password:           c.DB.Password,
		MaxIdleConnCount:   c.DB.MaxIdleConnections,
		MaxOpenConnCount:   c.DB.MaxOpenConnections,
		ConnMaxIdleTimeSec: c.DB.ConnMaxIdleTimeSec,
		ConnMaxLifeTimeSec: c.DB.ConnMaxLifeTimeSec,
		Dialect:            DialectPostgres,
		Options: &gorm.Config{
			Logger:                 logger.Default.LogMode(logger.Info),
			SkipDefaultTransaction: true,
		},
	}

	for i := 0; i < ConnectRetryAttempts; i++ {
		db, err := dialDb(dbConfig)
		if err != nil {
			time.Sleep(RepositoryPostgresConnectWaitSec * time.Second)
			continue
		} else {
			return db, nil
		}
	}

	return nil, errors.New("Unable to connect to postgres after 10 attempts")
}

func dialDb(c DBConfig) (*gorm.DB, error) {
	dialector, serr := getDialector(c)
	if serr != nil {
		return nil, serr
	}

	dbObj, err := gorm.Open(
		dialector,
		c.Options,
	)
	if err != nil {
		return nil, errors.New("Unable to connect to postgres")
	}

	if db, err := dbObj.DB(); err == nil {
		// pooling setup
		if c.MaxIdleConnCount > 0 {
			db.SetMaxIdleConns(c.MaxIdleConnCount)
		}

		if c.MaxOpenConnCount > 0 {
			db.SetMaxOpenConns(c.MaxOpenConnCount)
		}

		// conn timeout setup
		if c.ConnMaxIdleTimeSec > 0 {
			db.SetConnMaxIdleTime(time.Duration(c.ConnMaxIdleTimeSec) * time.Second)
		}

		if c.ConnMaxLifeTimeSec > 0 {
			db.SetConnMaxLifetime(time.Duration(c.ConnMaxLifeTimeSec) * time.Second)
		}
	} else {
		return nil, errors.New("Unable to connect to postgres")
	}

	return dbObj, nil
}

func getDialector(c DBConfig) (gorm.Dialector, error) {
	switch c.Dialect {
	case DialectPostgres:
		return getPGDialector(c), nil

	default:
		var dialector gorm.Dialector
		return dialector, errors.New("dialect not supported")
	}
}

func getPGDialector(c DBConfig) gorm.Dialector {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s application_name=%s",
		c.Host,
		c.Port,
		c.User,
		c.DBName,
		c.Password,
		c.AppName,
	)

	return postgres.Open(dsn)
}
