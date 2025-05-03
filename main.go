package main

import (
	"fmt"
	"gpt-service-go/client"
	"gpt-service-go/config"
	"gpt-service-go/handler"
	"gpt-service-go/middleware"
	"gpt-service-go/service"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	log.Print("starting server...")
	e := echo.New()
	
	defer func() {
        if r := recover(); r != nil {
            log.Printf("Recovered from panic: %v", r)
        }
    }()
	// Create instances
	javaClient := client.NewJavaClient(cfg.JavaBackendAuthURL)
	openAIService := service.NewOpenAIService(cfg.OpenAIAPIKey)
	chatHandler := handler.NewChatHandler(openAIService)

	// JWT Middleware
	e.Use(middleware.JWTValidator(javaClient))

	// Routes
	e.POST("/chat", chatHandler.HandleChat)

	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf(":%s", cfg.Port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down server...")
	// shutdown the server with a 10-second timeout
	// graceful shutdown
	e.Close()
	}