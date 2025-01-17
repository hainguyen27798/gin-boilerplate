package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/gin-boilerplate/global"
	"github.com/hainguyen27798/gin-boilerplate/internal/initialize"
	"github.com/hainguyen27798/gin-boilerplate/pkg/helpers"
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

	helpers.Must(server.Run(fmt.Sprintf(":%s", global.AppConfig.Server.Port)))
}
