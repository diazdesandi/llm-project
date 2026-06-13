package container

import (
	"github.com/diazdesandi/llm-project/backend/config"
	"github.com/diazdesandi/llm-project/backend/internal/auth"
	"github.com/diazdesandi/llm-project/backend/internal/model"
	"github.com/diazdesandi/llm-project/backend/pkg/logger"
	supa "github.com/nedpals/supabase-go"
	"go.uber.org/zap"
)

type Container struct {
	ModelHandler *model.Handler
	AuthHandler  *auth.Handler
	Logger       *zap.Logger
}

func NewAppContainer(cfg *config.Config) (*Container, error) {

	// Config
	l := logger.CreateLogger()
	defaultModel := "tinyllama"

	// Clients
	modelClient := model.NewClient(cfg.OllamaURL)
	authClient := supa.CreateClient(cfg.SupabaseURL, cfg.SupabaseKey)

	// Wrap supabase.Client to implement auth.Client interface if needed
	wrappedAuthClient := auth.NewClient(authClient)

	// Services
	modelService := model.NewService(modelClient, defaultModel)
	authService := auth.NewService(*wrappedAuthClient)

	// Handlers
	modelHandler := model.NewHandler(*modelService)
	authHandler := auth.NewHandler(*authService)

	return &Container{
		Logger:       l,
		ModelHandler: modelHandler,
		AuthHandler:  authHandler,
	}, nil
}
