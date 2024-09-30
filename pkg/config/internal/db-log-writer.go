// Package internal @Author hubo 2024/9/29 19:16:00
package internal

import (
	"fmt"
	config2 "github.com/Hakunari/supertools-go/pkg/config"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

type Writer struct {
	config config2.DataBaseConfig
	writer logger.Writer
}

func NewWriter(config config2.DataBaseConfig, writer logger.Writer) *Writer {
	return &Writer{config, writer}
}

func (w *Writer) Printf(message string, data ...any) {
	if w.config.UseLogger {
		switch w.config.GetLogLevel() {
		case logger.Silent:
			zap.L().Debug(fmt.Sprintf(message, data...))
		case logger.Error:
			zap.L().Error(fmt.Sprintf(message, data...))
		case logger.Warn:
			zap.L().Warn(fmt.Sprintf(message, data...))
		case logger.Info:
			zap.L().Info(fmt.Sprintf(message, data...))
		default:
			zap.L().Info(fmt.Sprintf(message, data...))
		}
		return
	}
	w.writer.Printf(message, data...)
}
