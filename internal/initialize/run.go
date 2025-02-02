package initialize

import (
	"github.com/hainguyen27798/gin-boilerplate/global"
)

func Run() {
	LoadConfig()
	InitLogger()

	s := InitServer()

	defer s.Stop()
	s.Run(global.AppConfig.Server.Port)
}
