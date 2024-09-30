// Package initialize @Author hubo 2024/9/30 14:08:00
package initialize

import (
	"fmt"
	"github.com/Hakunari/supertools-go/pkg/config"
	"github.com/Hakunari/supertools-go/pkg/consul"
	"github.com/Hakunari/supertools-go/pkg/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Initializer[T config.IAppConfig, S config.IServiceLocalConfig] struct {
	Logger      *zap.Logger
	AppConfig   *T
	LocalConfig *S
}

func NewInitializer[T config.IAppConfig, S config.IServiceLocalConfig]() *Initializer[T, S] {
	return &Initializer[T, S]{}
}

func (i *Initializer[T, S]) InitLocalConfig(configPath string) error {
	localCfg, err := config.LoadLocalConfig[S](configPath)
	if err != nil {
		return err
	}
	i.LocalConfig = localCfg
	return nil
}

func (i *Initializer[T, S]) InitLogger() {
	localConfig := (*i.LocalConfig).GetBaseConfig()
	i.Logger = logger.InitLogger(localConfig.Logger)
	zap.ReplaceGlobals(i.Logger)
}

func (i *Initializer[T, S]) LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		i.Logger.Error("Error loading .env file", zap.Error(err))
		return err
	}
	return nil
}

func (i *Initializer[T, S]) RegisterService() error {
	localConfig := (*i.LocalConfig).GetBaseConfig()
	if err := consul.RegisterService(localConfig); err != nil {
		i.Logger.Error("Failed to register service to consul", zap.Error(err))
		return err
	} else {
		i.Logger.Info(fmt.Sprintf("Register service to consul successfully, serviceId: %s", localConfig.Service.Id))
	}
	return nil
}

func (i *Initializer[T, S]) LoadAppConfig() error {
	localConfig := (*i.LocalConfig).GetBaseConfig()
	consulAddr := fmt.Sprintf("%s:%d", localConfig.Consul.Host, localConfig.Consul.Port)
	configData, err := consul.LoadCfgFromConsul[T](consulAddr, localConfig.Service.Name)
	if err != nil {
		i.Logger.Error("Failed to load models from consul", zap.Error(err))
		return err
	}
	i.AppConfig = configData
	return nil
}

func (i *Initializer[T, S]) DeregisterService() {
	localConfig := (*i.LocalConfig).GetBaseConfig()
	if err := consul.DeRegisterService(localConfig); err != nil {
		i.Logger.Error("Failed to deregister service from consul", zap.Error(err))
	} else {
		i.Logger.Info("Deregister service from consul successfully")
	}
}
