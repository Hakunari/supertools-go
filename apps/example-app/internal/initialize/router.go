// Package initialize @Author hubo 2024/9/27 14:45:00
package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExampleAppRoutes(router *gin.Engine) {
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, SuperReport!",
		})
	})
}
