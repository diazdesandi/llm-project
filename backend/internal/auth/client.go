package auth

import (
	"context"

	supa "github.com/nedpals/supabase-go"
)

// NewClient initializes Supabase client using env vars.
func NewClient(db *supa.Client) *Client {
	return &Client{db: db}
}

func (c *Client) GetUser(ctx context.Context, token string) (*supa.User, error) {
	user, err := c.db.Auth.User(ctx, token)
	if err != nil {
		return nil, err
	}
	return user, nil
}
