package global

import (
	"github.com/go-playground/validator/v10"
	"github.com/hainguyen27798/gin-boilerplate/internal/database"
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
	MongoDB   *database.MongoDBStrategy
	Validator *validator.Validate
)
