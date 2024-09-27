// Package initialize @Author hubo 2024/9/26 21:36:00
package initialize

import (
	"github.com/Hakunari/supertools-go/server/apps/sys-app/internal/global"
	"github.com/Hakunari/supertools-go/server/pkg/initialize"
)

func RunServer() {
	initialize.RunServer(global.GlbLocalConfig, global.GlbLogger, SysAppRoutes)
}
