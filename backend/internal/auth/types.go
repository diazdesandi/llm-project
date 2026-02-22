package auth

import (
	supa "github.com/nedpals/supabase-go"
)

type Handler struct {
	Service Service
}

type Service struct {
	client Client
}

type Client struct {
	db *supa.Client
}
