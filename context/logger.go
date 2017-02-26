package context

import (
	"github.com/uber-go/zap"
	"fmt"
	"context"
)

// SetLogger sets logger to context.
func SetLogger(ctx context.Context, logger zap.Logger) context.Context {
	return context.WithValue(ctx, contextKey("LOGGER"), logger)
}

// GetLogger sets logger from context.
func GetLogger(ctx context.Context) (zap.Logger, error) {
	l, ok := ctx.Value(contextKey("LOGGER")).(zap.Logger)
	if ok {
		return l, nil
	}

	return nil, fmt.Errorf("Failed to get LOGGER from context")
}

// MustGetLogger 確実にLoggerを取得
func MustGetLogger(ctx context.Context) zap.Logger {
	l, err := GetLogger(ctx)
	if err != nil {
		panic(err)
	}
	return l
}
