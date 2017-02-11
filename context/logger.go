package context

import (
	"github.com/uber-go/zap"	
	"fmt"
	"context"
)

func SetLogger(ctx context.Context, logger zap.Logger) context.Context {
	return context.WithValue(ctx, "LOGGER", logger)
}

func GetLogger(ctx context.Context) (zap.Logger, error) {
	l, ok := ctx.Value("LOGGER").(zap.Logger)
	if ok {
		return l, nil
	}

	return nil, fmt.Errorf("Failed to get LOGGER from context")
}

func MustGetLogger(ctx context.Context) zap.Logger {
	l, err := GetLogger(ctx)
	if err != nil {
		panic(err)
	}
	return l
}
