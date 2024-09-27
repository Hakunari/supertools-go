// Package initialize @Author hubo 2024/9/26 16:04:00
package initialize

import (
	"github.com/Hakunari/supertools-go/server/apps/sys-app/internal/global"
	"github.com/Hakunari/supertools-go/server/apps/sys-app/internal/models"
	commonInit "github.com/Hakunari/supertools-go/server/pkg/initialize"
	"go.uber.org/zap"
)

func InitBase() {
	var err error
	global.GlbLogger, global.GlbAppConfig, err = commonInit.InitBase[models.AppConfig](&global.GlbLocalConfig)
	if err != nil {
		zap.L().Fatal("Initialization failed", zap.Error(err))
	}

}
