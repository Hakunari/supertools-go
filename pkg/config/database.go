// Package config @Author hubo 2024/9/27 17:07:00
package config

import (
	"fmt"
	"github.com/Hakunari/supertools-go/pkg/config/internal"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"net/url"
	"os"
	"strings"
	"time"
)

type GenDBConfig struct {
	Prefix   string `mapstructure:"prefix"`
	Singular bool   `mapstructure:"singular"`
	Engine   string `mapstructure:"engine"`
}

type DataBaseConfig struct {
	Driver       string                 `mapstructure:"driver"`
	Host         string                 `mapstructure:"host"`
	Port         int                    `mapstructure:"port"`
	User         string                 `mapstructure:"user"`
	Password     string                 `mapstructure:"password"`
	DbName       string                 `mapstructure:"dbname"`
	UseLogger    bool                   `mapstructure:"useLogger"`
	LogLevel     string                 `mapstructure:"logLevel"`
	GenConfig    GenDBConfig            `mapstructure:"gen"`
	MaxIdleConns int                    `mapstructure:"maxIdleConns"`
	MaxOpenConns int                    `mapstructure:"maxOpenConns"`
	ExtraKeys    []string               `mapstructure:"extraKeys"` // 由于 Viper 存在将 key 强制转为小写的问题(https://github.com/spf13/viper/issues/373), 需存储 ExtraParams 的 keys
	ExtraParams  map[string]interface{} `mapstructure:"extraParams"`
}

func (c *DataBaseConfig) GetLogLevel() logger.LogLevel {
	switch strings.ToLower(c.LogLevel) {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Info
	}
}

func (c *DataBaseConfig) GetExParamsUrl() string {
	params := url.Values{}
	for _, key := range c.ExtraKeys {
		lowerKey := strings.ToLower(key)
		params.Add(key, fmt.Sprintf("%v", c.ExtraParams[lowerKey]))
	}
	//for key, value := range c.ExtraParams {
	//	params.Add(key, fmt.Sprintf("%v", value))
	//}
	return params.Encode()
}

func (c *DataBaseConfig) GetGormConfig() *gorm.Config {
	return &gorm.Config{
		Logger: logger.New(internal.NewWriter(*c, log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      c.GetLogLevel(),
			Colorful:      true,
		}),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.GenConfig.Prefix,
			SingularTable: c.GenConfig.Singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}
