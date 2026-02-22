package auth

import (
	"context"
	"errors"
)

func (s *Service) SignIn(ctx context.Context, email, password string) (string, error) {

	details, err := s.client.SignIn(ctx, email, password)
	if err != nil {
		return "", errors.New("invalid email and/or password")
	}

	// We do not implement password hashing since Supabase handles that internally.

	// Return the access token (JWT) upon successful sign-in.
	if details.AccessToken == "" {
		return "", errors.New("failed to retrieve access token")
	}

	return details.AccessToken, nil

}
