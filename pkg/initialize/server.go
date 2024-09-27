// Package initialize @Author hubo 2024/9/27 14:50:00
package initialize

import (
	"fmt"
	"github.com/Hakunari/supertools-go/server/pkg/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer(localConfig *models.ServiceLocalConfig, logger *zap.Logger, customRoutes func(router *gin.Engine)) {
	router := InitRouters(customRoutes)

	port := fmt.Sprintf(":%d", localConfig.Service.Port)
	serv := initServer(port, router)

	logger.Info(
		"server running on ",
		zap.String(
			"address",
			fmt.Sprintf("%s:%d", localConfig.Service.Address, localConfig.Service.Port),
		),
	)

	err := serv.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
	}
}
