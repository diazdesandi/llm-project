package model

import (
	"context"
	"errors"
)

func NewService(c Client, defaultModel string) *Service {
	return &Service{
		client:       c,
		defaultModel: defaultModel,
	}
}

func (s *Service) GetResponse(ctx context.Context, req *Request) (*Response, error) {
	if req.Prompt == "" {
		return nil, errors.New("prompt cannot be empty")
	}

	if req.Model == "" {
		req.Model = s.defaultModel
	}

	return s.client.GetResponse(ctx, req)
}
