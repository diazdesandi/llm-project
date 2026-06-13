package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/diazdesandi/llm-project/backend/config"
	"github.com/diazdesandi/llm-project/backend/internal/container"
	"github.com/diazdesandi/llm-project/backend/internal/routes"
	"github.com/diazdesandi/llm-project/backend/internal/shared/metrics"
	"go.uber.org/zap"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize metrics
	metrics.InitMetrics()

	// Initialize App Container
	c, err := container.NewAppContainer(cfg)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := c.Logger.Sync(); err != nil {
			c.Logger.Error("Failed to sync logger", zap.Error(err))
		}
	}()

	logger := c.Logger

	r := routes.SetupRoutes(*c.ModelHandler, *c.AuthHandler)

	server := &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        r,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Start server
	go func() {
		logger.Info("Starting server on port " + cfg.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Failed to start server:" + err.Error())
		}
	}()

	select {}
}
