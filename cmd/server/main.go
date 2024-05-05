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
	"reshub/internal/database"
	"reshub/internal/handler"
	"reshub/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init config(environment variables)
	cfg := config.NewConfig()

	// Init log
	logWriter := setupLogger(cfg)
	defer logWriter.Close()
	if cfg.Env == "production" {
		gin.DefaultWriter = io.MultiWriter(logWriter)
	}
	logger.NewSlog(logWriter)

	// Init db
	var d database.Closer
	switch cfg.DBConnection {
	case "postgresql":
	case "postgres":
	case "psql":
		d = database.NewPSQL(cfg)
	case "mysql":
	case "mariadb":
		d = database.NewSQL(cfg)
	default:
		log.Fatal("failed to retrive db connections")
	}
	defer d.Close()

	// Set handler
	pingHandler := handler.NewPingHandler()

	router := gin.Default()

	router.Static("/static", "./static")

	router.GET("/api/ping", pingHandler.Ping)
	router.POST("/api/pong", pingHandler.Pong)

	// Setup and run server
	startServer(cfg, router)
}

func startServer(cfg *config.Config, router *gin.Engine) {
	if cfg.HTTPPort == "" {
		log.Fatal("HTTP port not specified in the configuration")
	}

	// Setup and run server
	serverAddr := fmt.Sprintf("%s:%s", cfg.HTTPHost, cfg.HTTPPort)
	log.Printf("starting server on server address: %s\n", serverAddr)
	srv := &http.Server{
		Addr:    serverAddr,
		Handler: router,
	}
	go func() {
		// Service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to listen and serve: %v", err)
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
	if runtime.GOOS == "linux" && cfg.Env == "production" {
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
		file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf("failed to open log file: %v", err)
		}
		return file
	}
}
