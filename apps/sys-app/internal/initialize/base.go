// Package initialize @Author hubo 2024/9/26 16:04:00
package initialize

import (
	"github.com/Hakunari/supertools-go/apps/sys-app/internal/global"
	"github.com/Hakunari/supertools-go/apps/sys-app/internal/models"
	commonModel "github.com/Hakunari/supertools-go/pkg/config"
	commonInit "github.com/Hakunari/supertools-go/pkg/initialize"
	"go.uber.org/zap"
)

func InitBase() {
	var err error
	initRes, err := commonInit.InitBase[models.SysAppConfig, commonModel.ServiceLocalConfig]()
	if err != nil {
		zap.L().Fatal("Initialization failed", zap.Error(err))
	}
	global.GlbLogger = initRes.Logger
	global.GlbAppConfig = initRes.AppConfig
	global.GlbLocalConfig = initRes.LocalConfig

}
