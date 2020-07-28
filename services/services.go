package services

import (
	"fmt"

	"git.hifx.in/concurrency/config"
	"git.hifx.in/concurrency/container"
	"git.hifx.in/concurrency/core/logger"
	"github.com/jinzhu/gorm"
)

//Load all services
func Load() error {

	MysqlMaster, err := getMsSQLConnection(
		config.Current.MysqlMasterUsername,
		config.Current.MysqlMasterPassword,
		config.Current.MysqlMasterHost,
		config.Current.MysqlMasterPort,
		config.Current.MysqlMasterDB)

	if err != nil {
		logger.Error.Error("Error while setting dbWriter", err)
		// return err
	}
	container.SetDbWriter(MysqlMaster)

	return nil
}

//getMsSqlConnection returns a db connection
func getMsSQLConnection(username string, password string, host string, port int, dbname string) (*gorm.DB, error) {
	const mysqlConenctionString = "%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local"
	connectionString := fmt.Sprintf(mysqlConenctionString, username, password, host, port, dbname)
	db, err := gorm.Open("mysql", connectionString)

	return db, err
}
