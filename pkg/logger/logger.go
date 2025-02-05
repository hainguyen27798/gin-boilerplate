package logger

import (
	"fmt"
	"os"

	"github.com/hainguyen27798/gin-boilerplate/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Zap is a wrapper around zap.Logger
// Providing additional logging functionality and integration capabilities.
type Zap struct {
	*zap.Logger
}

// NewLogger creates and returns a new instance of
// the Zap logger configured with the provided settings and application mode.
func NewLogger(config setting.LoggerSettings, mode setting.AppMode, version string) *Zap {
	logLevel := config.Level
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}
	encoder := getEncodeLog()
	sync := getWriteSync(config)
	core := zapcore.NewCore(encoder, sync, level)

	return &Zap{
		zap.New(
			core, zap.AddCaller(),
			zap.AddStacktrace(zap.ErrorLevel)).Named(fmt.Sprintf("[%s-%s]",
			mode,
			version),
		),
	}
}

// getEncodeLog returns a zapcore.Encoder configured with a console encoder
// and custom timestamp, level, and caller formatting.
func getEncodeLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	// format timestamp
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// ts -> Time
	encoderConfig.TimeKey = "time"

	// info -> INFO
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// "caller": "cli/main.log.go:3"
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getWriteSync creates a zapcore.WriteSyncer that combines console output
// and file output based on the provided configuration.
func getWriteSync(config setting.LoggerSettings) zapcore.WriteSyncer {
	hook := lumberjack.Logger{
		Filename:   config.FileName,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}
	syncConsole := zapcore.AddSync(os.Stdout)
	return zapcore.NewMultiWriteSyncer(syncConsole, zapcore.AddSync(&hook))
}
