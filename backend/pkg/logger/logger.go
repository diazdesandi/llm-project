package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func CreateLogger() *zap.Logger {
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel), // Para producción
		Development:      false,
		Encoding:         "json",
		OutputPaths:      []string{"stdout", "/tmp/myapp.log"}, // Ruta real
		ErrorOutputPaths: []string{"stderr"},
		InitialFields: map[string]any{
			"service": "backend",
		},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "timestamp",
			LevelKey:      "level",
			NameKey:       "logger",
			CallerKey:     "caller",
			MessageKey:    "message",
			StacktraceKey: "stacktrace",
			EncodeLevel:   zapcore.LowercaseLevelEncoder,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return logger
}
