package clients

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/diazdesandi/llm-project/backend/internal/dto"
)

var endpoint = "http://ollama:11434"

// TODO: Refactor for env variable.
func OllamaClient(body *dto.OllamaRequest) (*dto.OllamaResponse, error) {

	// Marshal the body to JSON
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// Ollama endpoint POST request
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode the response into OllamaResponse
	var ollamaResp dto.OllamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		return nil, err
	}

	return &ollamaResp, nil
}
