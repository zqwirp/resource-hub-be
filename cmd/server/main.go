package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/syslog"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"reshub/config"
	"reshub/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()

	logWriter := setupLogger(cfg)
	defer logWriter.Close()

	// logger.NewSlog(logFile)

	// Set handler
	pingHandler := handler.NewPingHandler()

	ginLogger := ginSetupLogger(cfg)
	defer ginLogger.Close()

	router := gin.Default()
	router.GET("/ping", pingHandler.Ping)
	router.POST("/pong", pingHandler.Pong)

	// Setup and run server
	startServer(cfg, router)
}

func ginSetupLogger(cfg *config.Config) *os.File {
	logFilePath := fmt.Sprintf("%s/http.log", cfg.LogPath)

	gin.DisableConsoleColor()
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}

	gin.DefaultWriter = io.MultiWriter(file)
	return file
}

func startServer(cfg *config.Config, router *gin.Engine) {
	if cfg.HTTPPort == "" {
		log.Fatal("HTTP port not specified in the configuration")
	}

	// Log server start
	log.Printf("Starting server on port: %s\n", cfg.HTTPPort)

	// Setup and run server
	srv := &http.Server{
		Addr:    "127.0.0.1:" + cfg.HTTPPort,
		Handler: router,
	}
	go func() {
		// Service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to listen and serve: %v", err)
		}
	}()

	// Setup graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}

	<-ctx.Done()
	log.Println("Server exited gracefully")
}

func setupLogger(cfg *config.Config) io.WriteCloser {
	if runtime.GOOS == "linux" {
		sysLogger, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL0, "Resource Hub")
		if err != nil {
			log.Fatalf("failed to set up syslog: %v", err)
		}
		log.SetOutput(sysLogger)
		return sysLogger
	} else {
		err := os.MkdirAll(cfg.LogPath, 0755)
		if err != nil {
			log.Fatalf("failed to create log directory: %v", err)
		}

		logFilePath := fmt.Sprintf("%s/%s.log", cfg.LogPath, cfg.LogFileName)
		logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf("failed to open log file: %v", err)
		}
		return logFile
	}
}

// func setupLogger(cfg *config.Config) *os.File {
// 	err := os.MkdirAll(cfg.LogPath, 0755)
// 	if err != nil {
// 		log.Fatalf("failed to create log directory: %v", err)
// 	}

// 	logFilePath := fmt.Sprintf("%s/%s.log", cfg.LogPath, cfg.LogFileName)
// 	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
// 	if err != nil {
// 		log.Fatalf("failed to open log file: %v", err)
// 	}
// 	return logFile
// }
