package model

import (
	"github.com/diazdesandi/llm-project/backend/internal/shared/metrics"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

/*
	 "model":"tinyllama:latest",
	  "prompt":"Tell me an interesting fact about Tijuana, Mexico",
	  "stream":false
	}
*/

func NewHandler(service Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) Handler(c *gin.Context) {
	serviceName := "llm-backend"
	clientIP := c.ClientIP()
	host := c.Request.Host

	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		metrics.GostServiceHandlerErrors.With(prometheus.Labels{
			"service": serviceName,
			"client":  clientIP,
			"host":    host,
			"kind":    "bad_request",
		}).Inc()
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Service.GetResponse(c.Request.Context(), &req)
	if err != nil {
		metrics.GostServiceHandlerErrors.With(prometheus.Labels{
			"service": serviceName,
			"client":  clientIP,
			"host":    host,
			"kind":    "internal_error",
		}).Inc()
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"data":    resp,
	})
}
