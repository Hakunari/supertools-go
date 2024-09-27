// Package models @Author hubo 2024/9/26 17:09:00
package models

import "github.com/Hakunari/supertools-go/server/pkg/models"

type AppConfig struct {
	models.IAppConfig `mapstructure:"-"`
	DbConfig          models.DataBaseConfig `mapstructure:"database"`
}
