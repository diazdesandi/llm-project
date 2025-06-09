package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/diazdesandi/llm-project/backend/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

// Current code version
var version = "1.0.0"

// Define Prometheus metrics

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	},
	[]string{"method", "path", "status"},
)
var durationRequest = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Duration of HTTP requests in seconds",
		Buckets: prometheus.DefBuckets,
	},
	[]string{"method", "path", "status"},
)

var greetingRequests = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "greeting_requests_total",
	Help: "Número total de peticiones al endpoint /greeting/{name}",
})

type Body struct {
	Name string `path:"name" maxLength:"30" required:"true" example:"John Doe"`
}

type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

// CLI Options
type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"8080"`
}

func init() {
	prometheus.MustRegister(totalRequests, durationRequest, greetingRequests)
}

func main() {

	l := logger.CreateLogger()

	defer l.Sync()

	if err := godotenv.Load("./.env"); err != nil {
		l.Warn("Warning: Error loading .env file. Continuing with environment variables or defaults.")
	}

	// URL: // http://localhost:8080/health?name=John
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	databaseURL := os.Getenv("DATABASE_URL")
	ollamaURL := os.Getenv("OLLAMA_URL")
	if databaseURL == "" {
		databaseURL = "localhost:5432"
	}
	if ollamaURL == "" {
		databaseURL = "localhost:11434"
	}

	l.Info(port)
	l.Info(databaseURL)
	l.Info(ollamaURL)

	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {

		router := chi.NewMux()
		api := humachi.New(router, huma.DefaultConfig("Backend - LLM", version))

		huma.Get(api, "/greeting/{name}", func(ctx context.Context, input *Body) (*GreetingOutput, error) {

			greetingRequests.Inc() // <-- aquí se cuenta la invocación

			resp := &GreetingOutput{}
			resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
			l.Info(resp.Body.Message)
			return resp, nil
		})

		// http://backend:8080/metrics
		router.Handle("/metrics", promhttp.Handler())

		hooks.OnStart(func() {
			fmt.Printf(("Starting server on port %d...\n"), options.Port)
			http.ListenAndServe(fmt.Sprintf(":%d", options.Port), router)
		})
	})

	cli.Run()
}
