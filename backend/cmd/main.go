package main

import (
	"context"
	"fmt"
	"net/http"
	"os" // Added for converting status code to string

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/diazdesandi/llm-project/backend/internal/routes"
	"github.com/diazdesandi/llm-project/backend/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

// Current code version
var version = "1.0.0"

// Define Prometheus metrics
var registry = prometheus.NewRegistry()

var (
	// GOST service metrics
	gostServices = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gost_services",
			Help: "Number of services",
		},
		[]string{"kind"},
	)

	gostServiceRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gost_service_requests_total",
			Help: "Total number of requests",
		},
		[]string{"service", "client", "host"},
	)

	gostServiceRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "gost_service_request_duration_seconds",
			Help: "Request duration in seconds",
		},
		[]string{"service", "client", "host"},
	)

	gostServiceTransferInputBytes = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gost_service_transfer_input_bytes_total",
			Help: "Total input bytes transferred",
		},
		[]string{"service", "client", "host"},
	)

	gostServiceTransferOutputBytes = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gost_service_transfer_output_bytes_total",
			Help: "Total output bytes transferred",
		},
		[]string{"service", "client", "host"},
	)

	gostServiceHandlerErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gost_service_handler_errors_total",
			Help: "Total handler errors",
		},
		[]string{"service", "client", "host", "kind"},
	)

	gostServiceInputBytes = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gost_service_input_bytes_total",
			Help: "Total input bytes",
		},
		[]string{"service", "client", "host"},
	)

	gostServiceOutputBytes = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gost_service_output_bytes_total",
			Help: "Total output bytes",
		},
		[]string{"service", "client", "host"},
	)
)

// var greetingRequests = prometheus.NewCounter(prometheus.CounterOpts{
// 	Name: "greeting_requests_total",
// 	Help: "Total requests to endpoint /greeting/{name}",
// })

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

	gostServices.With(prometheus.Labels{"kind": "service"}).Set(1)

	// Register collectors with custom registry
	registry.MustRegister(
		collectors.NewBuildInfoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		gostServices,
		gostServiceRequests,
		gostServiceRequestDuration,
		gostServiceTransferInputBytes,
		gostServiceTransferOutputBytes,
		gostServiceHandlerErrors,
		gostServiceInputBytes,
		gostServiceOutputBytes,
	)
}

// CORS
// TODO: Refactor
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	l := logger.CreateLogger()

	defer l.Sync()

	if err := godotenv.Load("./.env"); err != nil {
		l.Warn("Warning: Error loading .env file. Continuing with environment variables or defaults.")
	}

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
		ollamaURL = "http://ollama:11434"
	}

	l.Info(port)
	l.Info(databaseURL)
	l.Info(ollamaURL)

	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {

		router := chi.NewMux()
		router.Use(corsMiddleware)
		api := humachi.New(router, huma.DefaultConfig("Backend - LLM", version))
		routes.RegisterOllamaRoutes(api)

		huma.Get(api, "/greeting/{name}", func(ctx context.Context, input *Body) (*GreetingOutput, error) {
			serviceName := "llm-backend"
			clientIP := "127.0.0.1"
			host := "localhost"

			gostServiceRequests.With(prometheus.Labels{
				"service": serviceName,
				"client":  clientIP,
				"host":    host,
			}).Inc()

			inputBytes := 100.0
			outputBytes := 200.0

			gostServiceInputBytes.With(prometheus.Labels{
				"service": serviceName,
				"client":  clientIP,
				"host":    host,
			}).Add(inputBytes)

			gostServiceOutputBytes.With(prometheus.Labels{
				"service": serviceName,
				"client":  clientIP,
				"host":    host,
			}).Add(outputBytes)

			gostServiceTransferInputBytes.With(prometheus.Labels{
				"service": serviceName,
				"client":  clientIP,
				"host":    host,
			}).Add(inputBytes)

			gostServiceTransferOutputBytes.With(prometheus.Labels{
				"service": serviceName,
				"client":  clientIP,
				"host":    host,
			}).Add(outputBytes)

			timer := prometheus.NewTimer(gostServiceRequestDuration.With(prometheus.Labels{
				"service": serviceName,
				"client":  clientIP,
				"host":    host,
			}))
			defer timer.ObserveDuration()

			resp := &GreetingOutput{}
			resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
			l.Info(resp.Body.Message)
			return resp, nil
		})

		// http://backend:8080/metrics
		router.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))

		hooks.OnStart(func() {
			fmt.Printf(("Starting server on port %d...\n"), options.Port)
			http.ListenAndServe(fmt.Sprintf(":%d", options.Port), router)
		})
	})

	cli.Run()
}
