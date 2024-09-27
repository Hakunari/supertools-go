// Package models @Author hubo 2024/9/26 16:27:00
package models

import (
	"go.uber.org/zap/zapcore"
	"time"
)

type LoggerConfig struct {
	Level         string `mapstructure:"level"`          // 日志打印级别
	Prefix        string `mapstructure:"prefix"`         // 日志前缀
	Format        string `mapstructure:"format"`         // 输出
	Directory     string `mapstructure:"directory"`      // 输出路径
	EncodeLevel   string `mapstructure:"encode-level"`   // 编码级别
	StacktraceKey string `mapstructure:"stacktrace-key"` // 栈名
	ShowLine      bool   `mapstructure:"show-line"`      // 显示行
	LogInConsole  bool   `mapstructure:"log-in-console"` // 是否输出到控制台
	RetentionDay  int    `mapstructure:"retention-day"`  // 日志保留天数
}

func (c *LoggerConfig) GetLevels() []zapcore.Level {
	levels := make([]zapcore.Level, 0, 10)
	level, err := zapcore.ParseLevel(c.Level)
	if err != nil {
		level = zapcore.DebugLevel
	}
	for ; level <= zapcore.FatalLevel; level++ {
		levels = append(levels, level)
	}
	return levels
}

func (c *LoggerConfig) Encoder() zapcore.Encoder {
	config := zapcore.EncoderConfig{
		TimeKey:       "time",
		NameKey:       "name",
		LevelKey:      "level",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: c.StacktraceKey,
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(c.Prefix + t.Format("2024-01-01 16:40:05.000"))
		},
		EncodeLevel:    c.LevelEncoder(),
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	if c.Format == "json" {
		return zapcore.NewJSONEncoder(config)
	}
	return zapcore.NewConsoleEncoder(config)

}

// LevelEncoder 根据 EncodeLevel 返回 zapcore.LevelEncoder
// Author [SliverHorn](https://github.com/SliverHorn)
func (c *LoggerConfig) LevelEncoder() zapcore.LevelEncoder {
	switch {
	case c.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		return zapcore.LowercaseLevelEncoder
	case c.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case c.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case c.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}
