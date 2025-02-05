package initialize

import (
	"fmt"

	"github.com/hainguyen27798/gin-boilerplate/global"
	"github.com/hainguyen27798/gin-boilerplate/internal/database"
	"go.uber.org/zap"
)

func InitDatabase() {
	mongoConfig := global.AppConfig.MongoDB
	connStr := fmt.Sprintf(
		"mongodb://%s:%s/?directConnection=%t",
		mongoConfig.Host,
		mongoConfig.Port,
		mongoConfig.DirectConnection,
	)

	strategy := &database.MongoDBStrategy{}

	// Create a DBContext with the selected strategy.
	dbContext := database.NewDBContext(strategy)
	conn, err := dbContext.Connect(connStr, database.DBOptions{
		Username:    mongoConfig.Username,
		Password:    mongoConfig.Password,
		DBName:      mongoConfig.Database,
		MaxPoolSize: mongoConfig.MaxPoolSize,
		EnableLog:   mongoConfig.EnableLog,
	})
	if err != nil {
		global.Logger.Error("init database fail", zap.Error(err))
		panic(err)
	}
	global.Logger.Info("init and connect database success")

	if mongoStrategy, ok := conn.(*database.MongoDBStrategy); ok {
		global.MongoDB = mongoStrategy
	}
}
