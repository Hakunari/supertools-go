// Package initialize @Author hubo 2024/9/26 16:04:00
package initialize

import (
	"github.com/Hakunari/supertools-go/apps/sys-app/internal/global"
	"github.com/Hakunari/supertools-go/apps/sys-app/internal/models"
	commonInit "github.com/Hakunari/supertools-go/pkg/initialize"
	commonModel "github.com/Hakunari/supertools-go/pkg/models"
	"go.uber.org/zap"
)

func InitBase() {
	var err error
	global.GlbLogger, global.GlbAppConfig, global.GlbLocalConfig, err = commonInit.InitBase[models.SysAppConfig, commonModel.ServiceLocalConfig]()
	if err != nil {
		zap.L().Fatal("Initialization failed", zap.Error(err))
	}

}
