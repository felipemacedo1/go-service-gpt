package main

import (
	"fmt"
	"gpt-service-go/client"
	"gpt-service-go/handler"
	customMiddleware "gpt-service-go/middleware"
	"gpt-service-go/config"
	"gpt-service-go/middleware"
	"gpt-service-go/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"github.com/sirupsen/logrus"

	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("Error loading config: %v", err)
	}

	// Configure logrus
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)

	logger.Info("starting server...")
	e := echo.New()

	defer func() {
		if r := recover(); r != nil {
			logger.Errorf("Recovered from panic: %v", r)
		}
	}()

	// Create instances
	javaClient := client.NewJavaClient(cfg.JavaBackendAuthURL)
	openAIService := service.NewOpenAIService(cfg.OpenAIAPIKey, logger)
	chatHandler := handler.NewChatHandler(openAIService, logger)

	// Middleware
	e.Use(customMiddleware.RateLimiter)
	
	chatGroup := e.Group("")
	chatGroup.Use(middleware.JWTValidator(javaClient))

	// Routes
	chatGroup.POST("/chat", chatHandler.HandleChat)

	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf(":%s", cfg.Port)); err != nil && err != http.ErrServerClosed {
			logger.Error("shutting down the server")
		}
	}()
	
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down server...")
	// shutdown the server with a 10-second timeout
	// graceful shutdown
	e.Close()
	}