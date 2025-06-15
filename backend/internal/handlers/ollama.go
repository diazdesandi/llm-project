package handlers

import (
	"context"

	"github.com/diazdesandi/llm-project/backend/internal/dto"
	"github.com/diazdesandi/llm-project/backend/internal/services"
	"github.com/diazdesandi/llm-project/backend/pkg/logger"
)

/*
	 "model":"tinyllama:latest",
	  "prompt":"Tell me an interesting fact about Tijuana, Mexico",
	  "stream":false
	}
*/

func OllamaHandler(ctx context.Context, input *dto.OllamaRequest) (*dto.OllamaResponse, error) {

	l := logger.CreateLogger()

	defer l.Sync()

	response, err := services.FetchOllamaResponse(input)
	if err != nil {
		return nil, err
	}

	return response, nil

}
