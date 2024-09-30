// Package sys_app @Author hubo 2024/9/26 11:22:00
package main

import (
	"github.com/Hakunari/supertools-go/apps/sys-app/internal/global"
	"github.com/Hakunari/supertools-go/apps/sys-app/internal/initialize"
	pkgInit "github.com/Hakunari/supertools-go/pkg/initialize"
)

func main() {
	// 初始化基础实例
	initialize.InitBase()

	pkgInit.InitGorm(global.GlbAppConfig.DbConfig)

	initialize.RunServer()
}
