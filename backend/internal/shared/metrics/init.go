package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

func InitMetrics() {

	GostServices.With(prometheus.Labels{"kind": "service"}).Set(1)

	// Register collectors with custom registry
	prometheus.MustRegister(
		GostServices,
		GostServiceRequests,
		GostServiceRequestDuration,
		GostServiceRequestsInFlight,
		GostServiceTransferInputBytes,
		GostServiceTransferOutputBytes,
		GostServiceHandlerErrors,
		GostServiceChainErrors,
		GostServiceInputBytes,
		GostServiceOutputBytes,
	)
}
