// Package global @Author hubo 2024/9/27 14:22:00
package global

import (
	serviceModel "github.com/Hakunari/supertools-go/apps/example-app/internal/models"
	"go.uber.org/zap"
)

var (
	GlbLogger      *zap.Logger
	GlbLocalConfig *serviceModel.AppLocalConfig
	GlbAppConfig   *serviceModel.ExampleAppConfig
)
