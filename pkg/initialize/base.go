// Package initialize @Author hubo 2024/9/27 14:33:00
package initialize

import (
	"fmt"
	"github.com/Hakunari/supertools-go/pkg/consul"
	"github.com/Hakunari/supertools-go/pkg/logger"
	"github.com/Hakunari/supertools-go/pkg/models"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

// InitBase 初始化基础信息：本地配置实例（localConfig）/日志实例/consul 服务注册/获取consul的服务配置
func InitBase[T models.IAppConfig](localConfig **models.ServiceLocalConfig) (*zap.Logger, *T, error) {

	// 初始化本地配置
	wd, _ := os.Getwd()
	configPath := filepath.Join(wd, "config.yaml")
	localCfg, err := models.LoadLocalConfig(configPath)
	if err != nil {
		return nil, nil, err
	}
	*localConfig = localCfg

	// 初始化日志实例
	globalLogger := logger.InitLogger((*localConfig).Logger)

	// 加载 .env 到环境变量
	if err = godotenv.Load(); err != nil {
		globalLogger.Error("Error loading .env file", zap.Error(err))
		return nil, nil, err
	}

	// 在 consul 注册服务
	if err = consul.RegisterService(*localConfig); err != nil {
		globalLogger.Error("Failed to register service to consul", zap.Error(err))
	}

	// 从 consul 获取服务配置信息
	consulAddr := fmt.Sprintf("%s:%d", (*localConfig).Consul.Host, (*localConfig).Consul.Port)
	configData, err := consul.LoadCfgFromConsul[T](consulAddr, (*localConfig).Service.Name)
	if err != nil {
		globalLogger.Error("Failed to load models from consul", zap.Error(err))
		return nil, nil, err
	}

	return globalLogger, configData, nil
}
