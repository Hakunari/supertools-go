// Package consul @Author hubo 2024/9/26 14:06:00 register service to consul
package consul

import (
	"bytes"
	"fmt"
	"github.com/Hakunari/supertools-go/pkg/config"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"os"
)

func genServiceId(serviceName, serviceHost string, port int) string {
	return fmt.Sprintf("%s-%s-%d-%s", serviceName, serviceHost, port, uuid.New().String())
}

// RegisterService 注册服务
func RegisterService(localConfig *config.ServiceLocalConfig) error {

	// register to consul.
	consulConfig := api.DefaultConfig()
	client, err := api.NewClient(consulConfig)
	if err != nil {
		return err
	}

	localConfig.Service.Id = genServiceId(localConfig.Service.Name, localConfig.Service.Host, localConfig.Service.Port)

	registration := &api.AgentServiceRegistration{
		ID:      localConfig.Service.Id,
		Name:    localConfig.Service.Name,
		Address: localConfig.Service.Host,
		Port:    localConfig.Service.Port,
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%d/health", localConfig.Service.Host, localConfig.Service.Port),
			Interval: localConfig.Service.Check.Interval,
			Timeout:  localConfig.Service.Check.Timeout,
		},
	}

	if err = client.Agent().ServiceRegister(registration); err != nil {
		return err
	}

	return nil
}

// LoadCfgFromConsul 从 consul 获取指定服务的配置
func LoadCfgFromConsul[T config.IAppConfig](consulAddr, serviceName string) (*T, error) {
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

func DeRegisterService(localConfig *config.ServiceLocalConfig) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", localConfig.Consul.Host, localConfig.Consul.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	return client.Agent().ServiceDeregister(localConfig.Service.Id)
}
