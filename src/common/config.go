package common

import (
	"bytes"
	"strings"

	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v1"
)

type Config struct {
	App AppConfig
	DB DBConfig
}

var c *Config

func NewConfig() (*Config, error) {
	if c != nil {
		return c, nil
	}

	defaultConfig := &Config{}
	if err := defaults.Set(defaultConfig); err != nil {
		return nil, err
	}

	c, err := LoadConfig(defaultConfig)
	if err != nil {
		return nil, err
	}

	return c, nil
}

type AppConfig struct {
	Name string `default:"my-servoce"`
	Port string `default:"8080"`
}

type DBConfig struct {
	Host               string `default:"localhost"`
	Port               int    `default:"5432"`
	Name               string `default:""`
	User               string `default:""`
	Password           string `default:""`
	MaxIdleConnections int    `default:"2"`
	MaxOpenConnections int    `default:"5"`
	ConnMaxIdleTimeSec int    `default:"10"`
	ConnMaxLifeTimeSec int    `default:"600"`
}

func LoadConfig[T any](c T) (T, error) {
	err := SetDefault(c)
	if err != nil {
		return c, err
	}

	bs, e := yaml.Marshal(c)
	if e != nil {
		return c, err
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if e = viper.ReadConfig(bytes.NewBuffer(bs)); e != nil {
		return c, err
	}

	bindKeys()

	e = viper.Unmarshal(c)
	if e != nil {
		return c, err
	}

	validate := validator.New()
	err = validate.Struct(c)
	if err != nil {
		return c, err
	}

	return c, nil
}

func bindKeys() {
	for _, key := range viper.AllKeys() {
		_ = viper.BindEnv(key, strings.ToUpper(strings.ReplaceAll(key, ".", "_")))
	}
}

func SetDefault[T any](c T) error {
	if err := defaults.Set(c); err != nil {
		return err
	}

	return nil
}
