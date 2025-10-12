package main

import (
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/diazdesandi/llm-project/backend/config"
	"github.com/diazdesandi/llm-project/backend/internal/metrics"
	"github.com/diazdesandi/llm-project/backend/internal/middlewares"
	"github.com/diazdesandi/llm-project/backend/internal/routes"
	"github.com/diazdesandi/llm-project/backend/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

// Current code version
var version = "1.0.0"

// CLI Options
// TODO: Check huma requires options
type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"8080"`
}

func main() {
	// Load configuration
	config := config.LoadConfig()
	l := logger.CreateLogger()
	defer l.Sync()

	// Initialize metrics
	metrics.InitMetrics()


	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {

		router := chi.NewMux()
    	router.Use(middlewares.CORS)
		api := humachi.New(router, huma.DefaultConfig("Backend - LLM", version))
		routes.RegisterRoutes(api)

		// http://backend:8080/metrics
    	router.Handle("/metrics", promhttp.HandlerFor(metrics.Registry, promhttp.HandlerOpts{}))

		hooks.OnStart(func() {
			fmt.Printf(("Starting server on port %s...\n"), config.Port)
			http.ListenAndServe(fmt.Sprintf(":%s", config.Port), router)
		})
	})

	cli.Run()
}
