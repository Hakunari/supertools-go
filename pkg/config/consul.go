// Package config @Author hubo 2024/9/26 17:05:00
package config

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
