// Package initialize @Author hubo 2024/9/27 14:50:00
package initialize

import (
	"fmt"
	"github.com/Hakunari/supertools-go/pkg/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer(localConfig models.IServiceLocalConfig, logger *zap.Logger, customRoutes func(router *gin.Engine)) {
	router := InitRouters(customRoutes)

	baseLocalConfig := localConfig.GetBaseConfig()

	port := fmt.Sprintf(":%d", baseLocalConfig.Service.Port)
	serv := initServer(port, router)

	logger.Info(
		"server running on ",
		zap.String(
			"address",
			fmt.Sprintf("%s:%d", baseLocalConfig.Service.Host, baseLocalConfig.Service.Port),
		),
	)

	err := serv.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
	}
}
