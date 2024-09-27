//go:build !windows
// +build !windows

// Package initialize @Author hubo 2024/9/27 15:08:00
package initialize

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

func initServer(address string, router *gin.Engine) server {
	serv := endless.NewServer(address, router)
	serv.ReadHeaderTimeout = 10 * time.Minute
	serv.WriteTimeout = 10 * time.Minute
	serv.MaxHeaderBytes = 1 << 20
	return serv
}
