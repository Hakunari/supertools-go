// Package initialize @Author hubo 2024/9/29 20:13:00
package initialize

import (
	"fmt"
	"github.com/Hakunari/supertools-go/pkg/config"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitSQLServerGorm(config config.DataBaseConfig) *gorm.DB {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&%s", config.User, config.Password, config.Host, config.Port, config.DbName, config.GetExParamsUrl())
	mysqlCfg := sqlserver.Config{
		DSN:               dsn,
		DefaultStringSize: 255,
	}
	db, err := gorm.Open(sqlserver.New(mysqlCfg), config.GetGormConfig())
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	return db
}
