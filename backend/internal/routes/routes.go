package routes

import (
	"context"
	"fmt"

	"github.com/danielgtaylor/huma/v2"
	"github.com/diazdesandi/llm-project/backend/internal/handlers"
	"github.com/diazdesandi/llm-project/backend/internal/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type Body struct {
    Name string `path:"name" maxLength:"30" required:"true" example:"John Doe"`
}

type GreetingOutput struct {
    Body struct {
        Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
    }
}


func RegisterRoutes(app huma.API) {
    huma.Get(app, "/greeting/{name}", func(ctx context.Context, input *Body) (*GreetingOutput, error) {
        serviceName := "llm-backend"
        clientIP := "127.0.0.1"
        host := "localhost"

		metrics.GostServiceRequests.With(prometheus.Labels{
			"service": serviceName,
			"client":  clientIP,
			"host":    host,
		}).Inc()

		timer := prometheus.NewTimer(metrics.GostServiceRequestDuration.With(prometheus.Labels{
			"service": serviceName,
			"client":  clientIP,
			"host":    host,
		}))
		defer timer.ObserveDuration()

		inputBytes := 100.0
		outputBytes := 200.0

		metrics.GostServiceInputBytes.With(prometheus.Labels{
			"service": serviceName,
			"client":  clientIP,
			"host":    host,
		}).Add(inputBytes)

		metrics.GostServiceOutputBytes.With(prometheus.Labels{
			"service": serviceName,
			"client":  clientIP,
			"host":    host,
		}).Add(outputBytes)

		metrics.GostServiceTransferInputBytes.With(prometheus.Labels{
			"service": serviceName,
			"client":  clientIP,
			"host":    host,
		}).Add(inputBytes)

		metrics.GostServiceTransferOutputBytes.With(prometheus.Labels{
			"service": serviceName,
			"client":  clientIP,
			"host":    host,
		}).Add(outputBytes)

        resp := &GreetingOutput{}
        resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
        return resp, nil
    })

	huma.Post(app, "/model", handlers.OllamaHandler)
	
	// Auth routes
	// huma.Post(app, "/register", handlers.RegisterHandler)
	// huma.Post(app, "/login", handlers.LoginHandler)
	// huma.Get(app, "/profile", handlers.ProfileHandler)
}