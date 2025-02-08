package initialize

import (
	"context"
	"time"

	"github.com/hainguyen27798/gin-boilerplate/internal/routes"

	"github.com/hainguyen27798/gin-boilerplate/global"
)

func Run() {
	LoadConfig("./configs/")
	InitLogger()
	InitDatabase()

	RegisterValidations()

	s := InitServer()

	// Register routes
	routes.RegisterRoutes(s.r)

	defer func() {
		// Create a context with timeout for graceful shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// shutdown server
		s.Stop(ctx)

		// disconnect mongoDB
		err := global.MongoDB.Disconnect(ctx)
		if err != nil {
			panic(err.Error())
		}
		global.Logger.Info("MongoDB disconnect success")

		// after server shutdown
		<-ctx.Done()
		global.Logger.Info("Server shutdown")
	}()
	s.Run(global.AppConfig.Server.Port)
}
