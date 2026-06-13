package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Reminder
// Go uses caps for public (exported) names and minuscules for private (internal) names.

// Define Prometheus metrics
var (
	// GostServices metrics
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

	GostServiceRequestsInFlight = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gost_service_requests_in_flight",
			Help: "Current requests in flight",
		},
		[]string{"service", "client", "host"},
	)

	GostServiceChainErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gost_service_chain_errors_total",
			Help: "Total chain errors",
		},
		[]string{"service", "client", "host", "chain"},
	)
)
