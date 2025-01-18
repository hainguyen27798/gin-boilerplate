package initialize

import (
	"github.com/hainguyen27798/gin-boilerplate/global"
	"github.com/hainguyen27798/gin-boilerplate/pkg/logger"
	"github.com/hainguyen27798/gin-boilerplate/version"
)

// InitLogger initializes the global logger instance using application configuration settings for structured logging.
func InitLogger() {
	global.Logger = logger.NewLogger(global.AppConfig.Logger, global.AppMode, version.Version)
}
