package global

import "github.com/hainguyen27798/gin-boilerplate/pkg/setting"

// AppConfig holds the application's server configuration settings from the setting.Config structure.
// AppMode specifies the current operating mode of the application using the setting.AppMode type.
var (
	AppConfig setting.Config
	AppMode   setting.AppMode
)
