// Package initialize @Author hubo 2024/9/27 14:33:00
package initialize

import (
	"github.com/Hakunari/supertools-go/pkg/config"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

type InitResult[T config.IAppConfig, S config.IServiceLocalConfig] struct {
	Logger            *zap.Logger
	AppConfig         *T
	LocalConfig       *S
	DeregisterService func()
}

// InitBase 初始化基础信息：本地配置实例（localConfig）/日志实例/consul 服务注册/获取consul的服务配置
func InitBase[T config.IAppConfig, S config.IServiceLocalConfig]() (*InitResult[T, S], error) {

	initializer := NewInitializer[T, S]()

	wd, _ := os.Getwd()
	configPath := filepath.Join(wd, "config.yaml")

	if err := initializer.InitLocalConfig(configPath); err != nil {
		return nil, err
	}

	initializer.InitLogger()

	if err := initializer.LoadEnv(); err != nil {
		return nil, err
	}

	if err := initializer.RegisterService(); err != nil {
		return nil, err
	}

	if err := initializer.LoadAppConfig(); err != nil {
		return nil, err
	}

	deregister := func() {
		initializer.DeregisterService()
	}

	result := &InitResult[T, S]{
		Logger:            initializer.Logger,
		AppConfig:         initializer.AppConfig,
		LocalConfig:       initializer.LocalConfig,
		DeregisterService: deregister,
	}
	return result, nil

	//// 初始化本地配置
	//wd, _ := os.Getwd()
	//configPath := filepath.Join(wd, "config.yaml")
	//localCfg, err := config.LoadLocalConfig[S](configPath)
	//if err != nil {
	//	return nil, nil, nil, err
	//}
	//
	//localConfig := (*localCfg).GetBaseConfig()
	//
	//// 初始化日志实例
	//globalLogger := logger.InitLogger(localConfig.Logger)
	//
	//// 加载 .env 到环境变量
	//if err = godotenv.Load(); err != nil {
	//	globalLogger.Error("Error loading .env file", zap.Error(err))
	//	return nil, nil, nil, err
	//}
	//
	//// 在 consul 注册服务
	//if err = consul.RegisterService(localConfig); err != nil {
	//	globalLogger.Error("Failed to register service to consul", zap.Error(err))
	//}
	//
	//// 从 consul 获取服务配置信息
	//consulAddr := fmt.Sprintf("%s:%d", localConfig.Consul.Host, localConfig.Consul.Port)
	//configData, err := consul.LoadCfgFromConsul[T](consulAddr, localConfig.Service.Name)
	//if err != nil {
	//	globalLogger.Error("Failed to load models from consul", zap.Error(err))
	//	return nil, nil, nil, err
	//}
	//
	//return globalLogger, configData, localCfg, nil
}
