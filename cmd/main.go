package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/gin-boilerplate/configs"
	"github.com/hainguyen27798/gin-boilerplate/internal/initialize"
	"net/http"
)

func main() {
	initialize.Run()

	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err := server.Run(fmt.Sprintf(":%s", configs.AppConfig.Server.Port))
	if err != nil {
		panic(err)
	}
}
