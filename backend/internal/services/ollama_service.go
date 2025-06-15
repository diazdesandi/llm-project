package services

import (
	"github.com/diazdesandi/llm-project/backend/internal/clients"
	"github.com/diazdesandi/llm-project/backend/internal/dto"
)

func FetchOllamaResponse(body *dto.OllamaRequest) (*dto.OllamaResponse, error) {
	return clients.OllamaClient(body)
}
