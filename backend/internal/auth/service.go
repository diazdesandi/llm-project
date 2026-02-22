package auth

import (
	"context"
	"errors"

	supa "github.com/nedpals/supabase-go"
)

func NewService(c Client) *Service {
	return &Service{
		client: c,
	}
}

func (s *Service) GetProfile(ctx context.Context, token string) (*supa.User, error) {
	authUser, err := s.client.GetUser(ctx, token)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// var firstName, lastName string
	// if fn, ok := authUser.UserMetadata["first_name"].(string); ok {
	// 	firstName = fn
	// }
	// if ln, ok := authUser.UserMetadata["last_name"].(string); ok {
	// 	lastName = ln
	// }

	profile := &supa.User{
		ID:        authUser.ID,
		Email:     authUser.Email,
		CreatedAt: authUser.CreatedAt,
		UpdatedAt: authUser.UpdatedAt,
	}

	return profile, nil
}
