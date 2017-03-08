package middleware

import (
	"net/http"
	"go.uber.org/zap"
	"github.com/shinofara/simple-go-web-app/http/context"
)

// LoggerMiddleware loggerをwrap
type LoggerMiddleware struct {
	logger *zap.SugaredLogger
}

// NewLoggerMiddleware creates a loggerMiddleware
func NewLoggerMiddleware() *LoggerMiddleware {
	logger, _ := zap.NewProduction()
	
	return &LoggerMiddleware{
		logger: logger.Sugar(),
	}
}

// LoggerMiddleware stores Logger to context.
func (ml *LoggerMiddleware) LoggerMiddleware(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.SetLogger(r.Context(), ml.logger)
		r = r.WithContext(ctx)
		
		ml.logger.Info("Set logger to context.")
			next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
