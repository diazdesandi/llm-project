package model

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func NewClient(url string) *ClientContent {
	return &ClientContent{
		endpoint: url,
	}
}

func (c *ClientContent) GetResponse(ctx context.Context, body *Request) (*Response, error) {

	// Marshal the body to JSON
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Model endpoint POST request
	endpoint := c.endpoint + "/api/generate"

	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}

	defer response.Body.Close()

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Check HTTP status code
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("model API error %s: %s", response.Status, string(respBody))
	}

	// Decode for type checking
	var modelResp Response
	if err := json.Unmarshal(respBody, &modelResp); err != nil {
		return nil, fmt.Errorf("failed to decode model response: %w", err)
	}

	return &modelResp, nil
}
