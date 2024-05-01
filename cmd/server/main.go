package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"reshub/config"
	"reshub/internal/handler"
	"reshub/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	// Setup config
	cfg := config.NewConfig()

	// Set logger
	logPath := cfg.LogPath
	logFile := cfg.LogFile
	logLogger, err := logger.New(logPath, logFile)
	if err != nil {
		panic(err)
	}
	defer logLogger.File.Close() // Close file at program exit

	// Setup gin logLogger
	ginLogger, err := ginLogLogger(logPath, "http")
	if err != nil {
		panic(err)
	}
	defer ginLogger.Close()

	// Set handler
	pingHandler := handler.NewPingHandler()

	// Set router
	router := gin.Default()
	// Set routes
	router.GET("/ping", pingHandler.Ping)
	router.POST("/pong", pingHandler.Pong)

	// Set HTTP port number
	httpPort := cfg.HTTPPort
	if httpPort == "" {
		msg := "failed to get http port environment variable"
		logLogger.Error(msg)
		panic(msg)
	}
	logLogger.Info(fmt.Sprintf("starting server on port: %s", cfg.HTTPPort))

	// Setup and run server
	srv := &http.Server{
		Addr:    "127.0.0.1:" + cfg.HTTPPort,
		Handler: router,
	}
	go func() {
		// Service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logLogger.Error(err.Error())
		}
	}()

	// Shutdown program in a gracefully way...
	gracefullyShutDownServer(srv, logLogger)
}

func ginLogLogger(path, name string) (*os.File, error) {
	gin.DisableConsoleColor()
	file, err := os.OpenFile("./"+path+"/"+name+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	gin.DefaultWriter = io.MultiWriter(file)
	return file, nil
}

func gracefullyShutDownServer(srv *http.Server, l *logger.Logger) {
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)

	// kill (no param) default send syscanll.SIGTERM, -2 is syscall.SIGINT,
	// -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	l.Info("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		l.Info("server shutdown:", err)
	}

	<-ctx.Done()
	l.Info("server exiting")
}
