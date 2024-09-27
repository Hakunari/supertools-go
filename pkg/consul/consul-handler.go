// Package consul @Author hubo 2024/9/26 14:06:00 register service to consul
package consul

import (
	"bytes"
	"fmt"
	"github.com/Hakunari/supertools-go/pkg/models"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"os"
)

func RegisterService(glbLocalConfig *models.ServiceLocalConfig) error {

	// register to consul.
	consulConfig := api.DefaultConfig()
	client, err := api.NewClient(consulConfig)
	if err != nil {
		return err
	}

	registration := &api.AgentServiceRegistration{
		Name:    glbLocalConfig.Service.Name,
		Address: glbLocalConfig.Service.Address,
		Port:    glbLocalConfig.Service.Port,
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%d/health", glbLocalConfig.Service.Address, glbLocalConfig.Service.Port),
			Interval: glbLocalConfig.Service.Check.Interval,
			Timeout:  glbLocalConfig.Service.Check.Timeout,
		},
	}

	if err = client.Agent().ServiceRegister(registration); err != nil {
		return err
	}

	return nil
}

// LoadCfgFromConsul
//
//	@Description: 从 consul 获取指定服务的配置
//	@param consulAddr Consul 地址
//	@param serviceName 服务名称
//	@return models
//	@return err
func LoadCfgFromConsul[T models.IAppConfig](consulAddr, serviceName string) (*T, error) {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = consulAddr
	client, err := api.NewClient(consulConfig)
	if err != nil {
		return nil, err
	}

	// 获取环境变量 SR_PROFILE
	profile := os.Getenv("SR_PROFILE")
	if profile == "" {
		profile = "dev"
	}
	key := fmt.Sprintf("config/%s-%s.yaml", serviceName, profile)
	kv := client.KV()
	pair, _, err := kv.Get(key, nil)
	if err != nil {
		return nil, err
	} else if pair == nil {
		return nil, nil
	}

	// 将配置从 byte[] 转为 T
	v := viper.New()
	v.SetConfigType("yaml")
	if err = v.ReadConfig(bytes.NewBuffer(pair.Value)); err != nil {
		return nil, err
	}

	var cfg T
	if err = v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, err
}
