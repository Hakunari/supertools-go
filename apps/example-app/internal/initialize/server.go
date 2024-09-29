// Package initialize @Author hubo 2024/9/27 15:12:00
package initialize

import (
	"github.com/Hakunari/supertools-go/apps/example-app/internal/global"
	"github.com/Hakunari/supertools-go/pkg/initialize"
	"github.com/Hakunari/supertools-go/pkg/models"
)

func RunServer() {
	var config models.IServiceLocalConfig = global.GlbLocalConfig
	initialize.RunServer(config, global.GlbLogger, ExampleAppRoutes)
}
