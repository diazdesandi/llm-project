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

var endpoint = "http://ollama:11434/api/generate" // Corrected Ollama API endpoint

func getDefaultModel() string {
	if model := os.Getenv("OLLAMA_MODEL"); model != "" {
		return model
	}
	return "tinyllama"
}

// TODO: Refactor for env variable.
func OllamaClient(body *dto.OllamaRequest) (*dto.OllamaResponse, error) {

	// Validate model
	if body.Model == "" {
		body.Model = getDefaultModel()
	}

	// Marshal the body to JSON
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Ollama endpoint POST request
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		// Attempt to read error body for more details, if any
		errorBodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("ollama API request failed with status %s: %s", resp.Status, string(errorBodyBytes))
	}

	// Decode the response into OllamaResponse
	var ollamaResp dto.OllamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		return nil, fmt.Errorf("failed to decode ollama response: %w", err)
	}

	return &ollamaResp, nil
}
