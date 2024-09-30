// Package models @Author hubo 2024/9/26 17:09:00
package models

import (
	"github.com/Hakunari/supertools-go/pkg/config"
)

// SysAppConfig sys-app 服务 KV 配置结构体
type SysAppConfig struct {
	config.AppConfig `mapstructure:",squash"`
	TestStr          string `mapstructure:"test-str"`
}

// GetBaseConfig 获取基础 KV 配置
func (s SysAppConfig) GetBaseConfig() *config.AppConfig {
	return &s.AppConfig
}
