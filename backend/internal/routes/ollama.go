package routes

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/diazdesandi/llm-project/backend/internal/handlers"
)

func RegisterOllamaRoutes(app huma.API) {
	huma.Post(app, "/model", handlers.OllamaHandler)
}
