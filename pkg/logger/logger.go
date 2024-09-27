// Package logger @Author hubo 2024/9/26 14:50:00
package logger

import (
	"fmt"
	"github.com/Hakunari/supertools-go/server/pkg/models"
	"github.com/Hakunari/supertools-go/server/pkg/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitLogger(loggerConfig models.LoggerConfig) (logger *zap.Logger) {
	if ok := utils.DirExists(loggerConfig.Directory); !ok {
		fmt.Printf("creating %v directory\n", loggerConfig.Directory)
		_ = os.Mkdir(loggerConfig.Directory, os.ModePerm)
	}
	levels := loggerConfig.GetLevels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := NewZapCore(levels[i], &loggerConfig)
		cores = append(cores, core)
	}
	logger = zap.New(zapcore.NewTee(cores...))
	if loggerConfig.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
