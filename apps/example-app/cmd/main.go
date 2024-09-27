// Package example_app @Author hubo 2024/9/27 15:15:00
package main

import (
	"github.com/Hakunari/supertools-go/apps/example-app/internal/initialize"
)

func main() {
	initialize.InitBase()

	initialize.RunServer()
}
