// Package sys_app @Author hubo 2024/9/26 11:22:00
package main

import (
	"github.com/Hakunari/supertools-go/apps/sys-app/internal/initialize"
)

func main() {
	// 初始化基础实例
	initialize.InitBase()
	// TODO: 初始化 gorm

	initialize.RunServer()
}
