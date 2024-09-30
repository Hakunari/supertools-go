// Package models @Author hubo 2024/9/27 14:21:00
package models

import (
	"github.com/Hakunari/supertools-go/pkg/config"
)

// ExampleAppConfig 示例服务 KV 配置结构体
type ExampleAppConfig struct {
	config.AppConfig `mapstructure:",squash"`
	TestStr          string `mapstructure:"test-str"`
}

// AppLocalConfig 示例服务本地配置结构体
type AppLocalConfig struct {
	config.ServiceLocalConfig `mapstructure:",squash"`
	TestStr                   string `mapstructure:"test-str"`
}

// GetBaseConfig 获取基础 KV 配置
func (s AppLocalConfig) GetBaseConfig() *config.ServiceLocalConfig {
	return &s.ServiceLocalConfig
}
