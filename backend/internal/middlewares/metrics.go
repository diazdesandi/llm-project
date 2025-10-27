package middlewares

import (
	"time"

	"github.com/diazdesandi/llm-project/backend/internal/shared/metrics"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		labels := prometheus.Labels{
			"service": "llm-backend",
			"client":  c.ClientIP(),
			"host":    c.Request.Host,
		}

		// Get bytes from request context
		inputBytes := float64(c.Request.ContentLength)

		metrics.GostServiceRequestsInFlight.With(labels).Inc()

		defer func() {

			// Gin native function to get response size
			outputBytes := float64(c.Writer.Size())

			metrics.GostServiceRequestsInFlight.With(labels).Dec()

			duration := time.Since(start).Seconds()
			metrics.GostServiceRequestDuration.With(labels).Observe(duration)

			metrics.GostServiceRequests.With(labels).Inc()

			metrics.GostServiceInputBytes.With(labels).Add(inputBytes)
			metrics.GostServiceOutputBytes.With(labels).Add(outputBytes)
			metrics.GostServiceTransferInputBytes.With(labels).Add(inputBytes)
			metrics.GostServiceTransferOutputBytes.With(labels).Add(outputBytes)
		}()

		c.Next()
	}
}
