package initialize

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/gin-boilerplate/global"
	"github.com/hainguyen27798/gin-boilerplate/pkg/setting"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	r *gin.Engine
	s *http.Server
}

// InitServer init gin server
func InitServer() *Server {
	var r *gin.Engine

	if global.AppMode == setting.ProdMode {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(gin.Logger(), gin.Recovery())
	} else {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	}

	return &Server{r, nil}
}

// Run server
func (s *Server) Run(port string) {
	// Create an HTTP server with the Gin router
	s.s = &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: s.r.Handler(),
	}

	// Start the server in a separate goroutine
	go func() {
		global.Logger.Info("Starting server on :" + port)

		// service connections
		if err := s.s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			global.Logger.Fatal("Server listen failed: \n" + err.Error())
			panic(err)
		}
	}()
}

func (s *Server) Stop(ctx context.Context) {
	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.Logger.Info("Shutting down server...")

	// Attempt to shut down the server gracefully
	if err := s.s.Shutdown(ctx); err != nil {
		global.Logger.Fatal("Server forced to shut down: \n" + err.Error())
		panic(err)
	}
}
