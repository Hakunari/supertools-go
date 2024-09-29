// Package initialize @Author hubo 2024/9/27 14:24:00
package initialize

import (
	"github.com/Hakunari/supertools-go/apps/example-app/internal/global"
	"github.com/Hakunari/supertools-go/apps/example-app/internal/models"
	commonInit "github.com/Hakunari/supertools-go/pkg/initialize"
	"go.uber.org/zap"
)

func InitBase() {
	var err error
	global.GlbLogger, global.GlbAppConfig, global.GlbLocalConfig, err = commonInit.InitBase[models.ExampleAppConfig, models.AppLocalConfig]()
	if err != nil {
		zap.L().Fatal("Initialization failed", zap.Error(err))
	}
}
