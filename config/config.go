package config

import (
	"git.hifx.in/concurrency/core/logger"
	"github.com/caarlos0/env"
)

//Config holds the env values
type Config struct {
	MysqlMasterHost     string
	MysqlMasterPort     int
	MysqlMasterDB       string
	MysqlMasterUsername string
	MysqlMasterPassword string
	LogFile             string
}

//Current env values
var Current Config

func Load() {

	Current.MysqlMasterPassword = "1234"
	Current.MysqlMasterUsername = "root"
	Current.MysqlMasterDB = "test"

	err := env.Parse(&Current)
	if err != nil {
		logger.Error.Error("Error while loading envs", err)

	}
}
