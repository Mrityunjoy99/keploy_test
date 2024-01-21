package migration

import (
	"fmt"

	"github.com/sample_project/cmd/cmdopts"
	"github.com/sample_project/src/common"
	"github.com/sample_project/src/infrastructure/database"

	"github.com/go-gormigrate/gormigrate/v2"
)

// nolint
func GetMigrations() []*gormigrate.Migration {
	migrations := []*gormigrate.Migration{
		V20240121034350,
	}

	return migrations
}

func Migrate(a *cmdopts.Arguments) {
	c, err := common.NewConfig()
	if err != nil {
		panic(err)
	}

	db, err := database.Connect(c)
	if err != nil {
		panic(err)
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, GetMigrations())

	switch {
	case cmdopts.IsFlagPassed(cmdopts.MigrateAllOpts.ToString()):
		err = m.Migrate()
	case cmdopts.IsFlagPassed(cmdopts.MigrateToOpts.ToString()):
		version := a.MigrateTo
		err = m.MigrateTo(version)
	case cmdopts.IsFlagPassed(cmdopts.MigrateLastOpts.ToString()):
		err = m.RollbackLast()
	case cmdopts.IsFlagPassed(cmdopts.RollbackToOpts.ToString()):
		version := a.RollbackTo
		err = m.RollbackTo(version)
	default:
		panic("Migration argument is a mandatory field.")
	}

	if err != nil {
		fmt.Printf("\nFailed to migrate: %v", err)
	}

	fmt.Println("Migation successful")
}
