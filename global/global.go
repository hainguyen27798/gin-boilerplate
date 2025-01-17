package global

import (
	"github.com/hainguyen27798/gin-boilerplate/pkg/logger"
	"github.com/hainguyen27798/gin-boilerplate/pkg/setting"
)

// AppConfig holds the application's main configuration including server and logger settings.
// AppMode specifies whether the application is running in development or production mode.
// Logger is a structured and leveled logger instance for application log management.
var (
	AppConfig setting.Config
	AppMode   setting.AppMode
	Logger    *logger.Zap
)
