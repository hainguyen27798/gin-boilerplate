package initialize

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/gin-boilerplate/global"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	LoadConfig()
	InitLogger()

	initServer()
}

func initServer() {
	// Init server
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Create an HTTP server with the Gin router
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", global.AppConfig.Server.Port),
		Handler: r.Handler(),
	}

	// Start the server in a separate goroutine
	go func() {
		global.Logger.Info("Starting server on :8080")

		// service connections
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			global.Logger.Fatal("Server listen failed: \n" + err.Error())
			panic(err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.Logger.Info("Shutting down server...")

	// Create a context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to shut down the server gracefully
	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Fatal("Server forced to shut down: \n" + err.Error())
		panic(err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		global.Logger.Info("timeout of 5 seconds.")
	}
	global.Logger.Info("Server exiting")
}
