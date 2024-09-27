// Package models @Author hubo 2024/9/26 14:14:00 service models model
package models

import "github.com/spf13/viper"

// IAppConfig
// @Description: Consul 上服务配置实体
type IAppConfig interface{}

// ServiceLocalConfig
// @Description: 本地配置实体
type ServiceLocalConfig struct {
	Service struct {
		Name    string `mapstructure:"name"`
		Address string `mapstructure:"address"`
		Port    int    `mapstructure:"port"`
		Check   struct {
			Interval string `mapstructure:"interval"`
			Timeout  string `mapstructure:"timeout"`
		} `mapstructure:"check"`
	} `mapstructure:"service"`
	Logger LoggerConfig `mapstructure:"logger"`
	Consul ConsulConfig `mapstructure:"consul"`
}

type ServiceConfig struct {
}

// LoadLocalConfig
//
//	@Description: Load service models from yaml.
//	@param configPath
//	@return *ServiceLocalConfig
//	@return error
func LoadLocalConfig(configPath string) (*ServiceLocalConfig, error) {
	v := viper.New()

	v.SetConfigFile(configPath)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var config ServiceLocalConfig
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
