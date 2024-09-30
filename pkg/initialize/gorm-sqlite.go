// Package initialize @Author hubo 2024/9/29 20:12:00
package initialize

import (
	"fmt"
	"github.com/Hakunari/supertools-go/pkg/config"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitSQLiteGorm(config config.DataBaseConfig) *gorm.DB {
	dsn := fmt.Sprintf("file:%s?%s", config.DbName, config.GetExParamsUrl())
	db, err := gorm.Open(sqlite.Open(dsn), config.GetGormConfig())
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	return db
}
