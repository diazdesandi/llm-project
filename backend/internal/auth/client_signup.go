package auth

import (
	"context"

	supa "github.com/nedpals/supabase-go"
)

func (c *Client) Signup(ctx context.Context, email, password, firstName, lastName string) (*supa.User, error) {

	data := map[string]interface{}{
		"first_name": firstName,
		"last_name":  lastName,
	}

	credentials := supa.UserCredentials{
		Email:    email,
		Password: password,
		Data:     data,
	}

	user, err := c.db.Auth.SignUp(ctx, credentials)
	if err != nil {
		return nil, err
	}
	return user, nil
}
