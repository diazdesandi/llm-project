package main

import (
	"net/http"
	"time"

	"github.com/diazdesandi/llm-project/backend/config"
	"github.com/diazdesandi/llm-project/backend/internal/container"
	"github.com/diazdesandi/llm-project/backend/internal/routes"
	"github.com/diazdesandi/llm-project/backend/internal/shared/metrics"
)

func main() {
	// Load configuration
	config := config.LoadConfig()

	// Initialize metrics
	metrics.InitMetrics()

	// Initalize App Container
	container, err := container.NewAppContainer(config)
	if err != nil {
		panic(err)
	}

	defer container.Logger.Sync()

	logger := container.Logger

	r := routes.SetupRoutes(*container.ModelHandler, *container.AuthHandler)

	server := &http.Server{
		Addr:           ":" + config.Port,
		Handler:        r,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Start server
	go func() {
		logger.Info("Starting server on port " + config.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Failed to start server:" + err.Error())
		}
	}()

	select {}
}
