package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/gin-boilerplate/global"
	"github.com/hainguyen27798/gin-boilerplate/internal/routes"
	"github.com/hainguyen27798/gin-boilerplate/internal/wires"
)

func RegisterRoutes(r *gin.Engine) {
	userController := wires.InitializeUserModule(global.MongoDB.DB)
	routes.RegisterUserRoutes(r, userController)
}
