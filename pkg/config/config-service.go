// Package config @Author hubo 2024/9/26 14:14:00 service models model
package config

import (
	"github.com/spf13/viper"
	"reflect"
)

// IAppConfig Consul 上服务配置接口(空接口)
type IAppConfig interface {
	GetBaseConfig() *AppConfig
}

// AppConfig Consul KV 存储基础配置
type AppConfig struct {
	DbConfig DataBaseConfig `mapstructure:"database"`
}

// GetBaseConfig 获取 KV 配置中的基础配置
func (a AppConfig) GetBaseConfig() *AppConfig {
	return &a
}

// IServiceLocalConfig 本地配置接口
type IServiceLocalConfig interface {
	GetBaseConfig() *ServiceLocalConfig
}

// ServiceLocalConfig 本地配置实体
type ServiceLocalConfig struct {
	Service struct {
		Id    string `mapstructure:"-"`
		Name  string `mapstructure:"name"`
		Host  string `mapstructure:"host"`
		Port  int    `mapstructure:"port"`
		Check struct {
			Interval string `mapstructure:"interval"`
			Timeout  string `mapstructure:"timeout"`
		} `mapstructure:"check"`
	} `mapstructure:"service"`
	Logger LoggerConfig `mapstructure:"logger"`
	Consul ConsulConfig `mapstructure:"consul"`
}

func (s ServiceLocalConfig) GetBaseConfig() *ServiceLocalConfig {
	return &s
}

// LoadLocalConfig Load service models from yaml.
func LoadLocalConfig[S IServiceLocalConfig](configPath string) (*S, error) {
	v := viper.New()

	v.SetConfigFile(configPath)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	// 需要通过反射创建 S 的实例, 否则会导致空接口反序列化失败, 所有属性为空.
	configType := reflect.TypeOf((*S)(nil)).Elem()
	configValue := reflect.New(configType).Interface()

	if err := v.Unmarshal(configValue); err != nil {
		return nil, err
	}

	result, _ := configValue.(*S)

	return result, nil
}
