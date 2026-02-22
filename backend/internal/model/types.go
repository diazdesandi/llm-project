package model

import "context"

type Client interface {
	GetResponse(ctx context.Context, req *Request) (*Response, error)
}

type ClientContent struct {
	endpoint string
}

type Handler struct {
	Service Service
}

type Service struct {
	client       Client
	defaultModel string
}
