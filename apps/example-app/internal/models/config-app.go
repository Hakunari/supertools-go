// Package models @Author hubo 2024/9/27 14:21:00
package models

import "github.com/Hakunari/supertools-go/pkg/models"

type ExampleAppConfig struct {
}

type AppLocalConfig struct {
	models.ServiceLocalConfig `mapstructure:",squash"`
	TestStr                   string `mapstructure:"test-str"`
}

func (s AppLocalConfig) GetBaseConfig() *models.ServiceLocalConfig {
	return &s.ServiceLocalConfig
}
