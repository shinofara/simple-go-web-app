package context

import (
	"go.uber.org/zap"
	"fmt"
	"context"
)

// SetLogger sets logger to context.
func SetLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, contextKey("LOGGER"), logger)
}

// GetLogger sets logger from context.
func GetLogger(ctx context.Context) (*zap.SugaredLogger, error) {
	l, ok := ctx.Value(contextKey("LOGGER")).(*zap.SugaredLogger)
	if ok {
		return l, nil
	}

	return nil, fmt.Errorf("Failed to get LOGGER from context")
}

// MustGetLogger 確実にLoggerを取得
func MustGetLogger(ctx context.Context) *zap.SugaredLogger {
	l, err := GetLogger(ctx)
	if err != nil {
		panic(err)
	}
	return l
}
