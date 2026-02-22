package auth

import (
	"context"
	"errors"
)

func (s *Service) Signup(ctx context.Context, email, password, firstName, lastName string) (string, error) {
	if len(password) < 8 {
		return "", errors.New("password must be at least 8 characters long")
	}

	user, err := s.client.Signup(ctx, email, password, firstName, lastName)
	if err != nil {
		return "", errors.New("failed to create user: " + err.Error())
	}

	return user.ID, nil
}
