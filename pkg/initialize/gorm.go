// Package initialize @Author hubo 2024/9/29 17:51:00
package initialize

import (
	"github.com/Hakunari/supertools-go/pkg/config"
	"github.com/Hakunari/supertools-go/pkg/constants"
	"gorm.io/gorm"
)

func InitGorm(config config.DataBaseConfig) *gorm.DB {
	switch config.Driver {
	case constants.MySQL.String():
		return InitMySQLGorm(config)
	case constants.PostgreSQL.String():
		return InitPGSQLGorm(config)
	case constants.SQLite.String():
		return InitSQLiteGorm(config)
	case constants.SQLServer.String():
		return InitSQLServerGorm(config)
	default:
		return InitMySQLGorm(config)
	}
}
