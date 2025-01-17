package logger

import (
	"github.com/hainguyen27798/gin-boilerplate/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// Zap is a wrapper around zap.Logger that provides structured and leveled logging capabilities.
type Zap struct {
	*zap.Logger
}

// NewLogger initializes and returns a new Zap logger instance based on the provided LoggerSettings configuration.
func NewLogger(config setting.LoggerSettings) *Zap {
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
		zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)),
	}
}

// getEncodeLog initializes and returns a JSON encoder for structured logging with custom time, level, and caller encodings.
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
	return zapcore.NewJSONEncoder(encoderConfig)
}

// getWriteSync creates a WriteSyncer that writes logs to both a file and the console based on the provided configuration.
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
