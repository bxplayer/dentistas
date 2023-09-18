package config

import (
	"errors"
	"os"
)

type Config struct {
	PublicConfig PublicConfig
}

type PublicConfig struct {
	MysqlHost     string
	MysqlPort     string
	MysqlUser     string
	MysqlPassword string
	MysqlDatabase string
}

var (
	envs = map[string]PublicConfig{
		"local": {
			MysqlHost:     os.Getenv("DB_HOST"),
			MysqlPort:     os.Getenv("DB_PORT"),
			MysqlUser:     os.Getenv("DB_USER"),
			MysqlPassword: os.Getenv("DB_PASS"),
			MysqlDatabase: os.Getenv("DB_NAME"),
		},
		"dev":  {},
		"prod": {},
	}
)

func NewConfig(env string) (Config, error) {

	publicConfig, exists := envs[env]
	if !exists {
		return Config{}, errors.New("env doest not exists")
	}

	return Config{
		PublicConfig: publicConfig,
	}, nil

}
