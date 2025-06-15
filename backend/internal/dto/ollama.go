package dto

import "time"

type OllamaRequest struct {
	Model  string `path:"model" maxLength:"30" required:"true" example:"tinyllama:latest"`
	Prompt string `path:"prompt" required:"true" example:"Tell me an interesting fact about Tijuana, Mexico"`
	Stream bool   `path:"stream" required:"true"`
}

type OllamaResponse struct {
	Model              string    `json:"model"`
	CreatedAt          time.Time `json:"created_at"`
	Response           string    `json:"response"`
	Done               bool      `json:"done"`
	DoneReason         string    `json:"done_reason"`
	Context            []int64   `json:"context"`
	TotalDuration      int64     `json:"total_duration"`
	LoadDuration       int64     `json:"load_duration"`
	PromptEvalCount    int64     `json:"prompt_eval_count"`
	PromptEvalDuration int64     `json:"prompt_eval_duration"`
	EvalCount          int64     `json:"eval_count"`
	EvalDuration       int64     `json:"eval_duration"`
}
