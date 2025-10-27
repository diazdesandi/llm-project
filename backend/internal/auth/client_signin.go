package auth

import (
	"context"

	supa "github.com/nedpals/supabase-go"
)

func (c *Client) SignIn(ctx context.Context, email, password string) (*supa.AuthenticatedDetails, error) {

	credentials := supa.UserCredentials{
		Email:    email,
		Password: password,
	}

	details, err := c.db.Auth.SignIn(ctx, credentials)
	if err != nil {
		return nil, err
	}
	return details, nil
}
