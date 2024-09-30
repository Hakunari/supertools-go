// Package initialize @Author hubo 2024/9/29 18:33:00
package initialize

import (
	"fmt"
	"github.com/Hakunari/supertools-go/pkg/config"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQLGorm(config config.DataBaseConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", config.User, config.Password, config.Host, config.Port, config.DbName, config.GetExParamsUrl())
	mysqlCfg := mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 255,
	}
	db, err := gorm.Open(mysql.New(mysqlCfg), config.GetGormConfig())
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	return db
}
