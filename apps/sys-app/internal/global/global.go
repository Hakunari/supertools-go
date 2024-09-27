// Package global @Author hubo 2024/9/26 15:00:00
package global

import (
	serviceCfg "github.com/Hakunari/supertools-go/apps/sys-app/internal/models"
	baseCfg "github.com/Hakunari/supertools-go/pkg/models"
	"go.uber.org/zap"
)

var (
	GlbLogger      *zap.Logger
	GlbLocalConfig *baseCfg.ServiceLocalConfig
	GlbAppConfig   *serviceCfg.AppConfig
)
