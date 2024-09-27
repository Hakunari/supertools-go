// Package initialize @Author hubo 2024/9/27 15:12:00
package initialize

import (
	"github.com/Hakunari/supertools-go/apps/example-app/internal/global"
	"github.com/Hakunari/supertools-go/pkg/initialize"
)

func RunServer() {
	initialize.RunServer(global.GlbLocalConfig, global.GlbLogger, ExampleAppRoutes)
}
