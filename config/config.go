package config

import (
	"git.hifx.in/release_manager/release_manager-API/core/logger"
	"github.com/caarlos0/env"
)

//Config holds the env values
type Config struct {
	MysqlMasterHost     string `env:"RELEASE_MANAGER_APP_API_MASTER_HOST,required"`
	MysqlMasterPort     int    `env:"RELEASE_MANAGER_APP_API_MASTER_PORT,required"`
	MysqlMasterDB       string `env:"RELEASE_MANAGER_APP_API_MASTER_DB,required"`
	MysqlMasterUsername string `env:"RELEASE_MANAGER_APP_API_MASTER_USERNAME,required"`
	MysqlMasterPassword string `env:"RELEASE_MANAGER_APP_API_MASTER_PASSWORD,required"`

	LogFile string `env:"RELEASE_MANAGER_APP_API_LOG_FILEPATH,required"`
}

//Current env values
var Current Config

//Load envs
func Load() {
	err := env.Parse(&Current)
	if err != nil {
		logger.Error.Error("Error while loading envs", err)

	}
}
