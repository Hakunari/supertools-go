// Package global @Author hubo 2024/9/27 14:22:00
package global

import (
	serviceModel "github.com/Hakunari/supertools-go/apps/example-app/internal/models"
	baseModel "github.com/Hakunari/supertools-go/pkg/models"
	"go.uber.org/zap"
)

var (
	GlbLogger      *zap.Logger
	GlbLocalConfig *baseModel.ServiceLocalConfig
	GlbAppConfig   *serviceModel.AppConfig
)
