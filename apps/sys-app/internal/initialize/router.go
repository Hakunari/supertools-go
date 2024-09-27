// Package initialize @Author hubo 2024/9/26 21:42:00
package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SysAppRoutes sys-app 服务路由初始化
func SysAppRoutes(router *gin.Engine) {
	// TODO: 后续在 controller 中添加 routers, 并在此处添加. 注: controller 根目录下应使用 groups.go 对路由进行分组
	// 测试路由 后删
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, SuperReport!",
		})
	})
}
