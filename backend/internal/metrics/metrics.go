package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

// Reminder
// Go uses mayuscules for public (exported) names and minuscules for private (internal) names.

// Define Prometheus metrics
var Registry = prometheus.NewRegistry()

var (
	// GOST service metrics
	GostServices = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gost_services",
			Help: "Number of services",
		},
		[]string{"kind"},
	)

	GostServiceRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gost_service_requests_total",
			Help: "Total number of requests",
		},
		[]string{"service", "client", "host"},
	)

	GostServiceRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "gost_service_request_duration_seconds",
			Help: "Request duration in seconds",
		},
		[]string{"service", "client", "host"},
	)

	GostServiceTransferInputBytes = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gost_service_transfer_input_bytes_total",
			Help: "Total input bytes transferred",
		},
		[]string{"service", "client", "host"},
	)

	GostServiceTransferOutputBytes = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gost_service_transfer_output_bytes_total",
			Help: "Total output bytes transferred",
		},
		[]string{"service", "client", "host"},
	)

	GostServiceHandlerErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gost_service_handler_errors_total",
			Help: "Total handler errors",
		},
		[]string{"service", "client", "host", "kind"},
	)

	GostServiceInputBytes = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gost_service_input_bytes_total",
			Help: "Total input bytes",
		},
		[]string{"service", "client", "host"},
	)

	GostServiceOutputBytes = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gost_service_output_bytes_total",
			Help: "Total output bytes",
		},
		[]string{"service", "client", "host"},
	)
)

func InitMetrics() {

	GostServices.With(prometheus.Labels{"kind": "service"}).Set(1)

	// Register collectors with custom registry
	Registry.MustRegister(
		collectors.NewBuildInfoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		GostServices,
		GostServiceRequests,
		GostServiceRequestDuration,
		GostServiceTransferInputBytes,
		GostServiceTransferOutputBytes,
		GostServiceHandlerErrors,
		GostServiceInputBytes,
		GostServiceOutputBytes,
	)
}