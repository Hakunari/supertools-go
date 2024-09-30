// Package initialize @Author hubo 2024/9/29 20:06:00
package initialize

import (
	"fmt"
	"github.com/Hakunari/supertools-go/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPGSQLGorm(config config.DataBaseConfig) *gorm.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?%s", config.User, config.Password, config.Host, config.Port, config.DbName, config.GetExParamsUrl())
	mysqlCfg := postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: false,
	}
	db, err := gorm.Open(postgres.New(mysqlCfg), config.GetGormConfig())
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	return db
}
