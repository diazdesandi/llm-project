package handlers

import (
	"context"

	"github.com/diazdesandi/llm-project/backend/internal/dto"
	"github.com/diazdesandi/llm-project/backend/internal/services"
)

/*
	 "model":"tinyllama:latest",
	  "prompt":"Tell me an interesting fact about Tijuana, Mexico",
	  "stream":false
	}
*/

type OllamaRequestWrapper struct {
	Body dto.OllamaRequest
}

type OllamaResponseWrapper struct {
	Body dto.OllamaResponse
}

func OllamaHandler(ctx context.Context, input *OllamaRequestWrapper) (*OllamaResponseWrapper, error) {

	response, err := services.FetchOllamaResponse(&input.Body)
	if err != nil {
		return nil, err
	}

	return &OllamaResponseWrapper{
		Body: *response,
	}, nil
}
