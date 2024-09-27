// Package initialize @Author hubo 2024/9/27 14:54:00
package initialize

import "github.com/gin-gonic/gin"

// InitCommonRouters 初始化服务通用路由
func InitCommonRouters(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.String(200, "ok")
	})
}

// InitRouters 接收一个自定义配置路由函数，完成总路由配置，并返回 gin.Engine
func InitRouters(customRoutes func(router *gin.Engine)) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		router.Use(gin.Logger())
	}

	// 初始化通用路由
	InitCommonRouters(router)

	// 调用自定义路由配置函数
	customRoutes(router)

	return router
}
