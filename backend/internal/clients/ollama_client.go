package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/diazdesandi/llm-project/backend/internal/dto"
)

// TODO: Refactor, implement main.go logger
func getOllamaEndpoint() string {
	if endpoint := os.Getenv("OLLAMA_URL"); endpoint != "" {
		return endpoint + "/api/generate"
	}
	return "http://ollama:11434/api/generate"
}

func getDefaultModel() string {
	if model := os.Getenv("OLLAMA_MODEL"); model != "" {
		return model
	}
	return "tinyllama"
}

func OllamaClient(body *dto.OllamaRequest) (*dto.OllamaResponse, error) {

	// Validate model
	if body.Model == "" {
		body.Model = getDefaultModel()
	}

	if body.Prompt == "" {
		return nil, fmt.Errorf("Prompt is empty")
	}

	// Marshal the body to JSON
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Ollama endpoint POST request
	endpoint := getOllamaEndpoint()
	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}

	defer response.Body.Close()

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Check HTTP status code
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ollama API error %s: %s", response.Status, string(respBody))
	}

	// Decode for type checking
	var ollamaResp dto.OllamaResponse
	if err := json.Unmarshal(respBody, &ollamaResp); err != nil {
		return nil, fmt.Errorf("failed to decode ollama response: %w", err)
	}

	return &ollamaResp, nil
}
