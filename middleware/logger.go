package middleware

import (
	"time"
	"net/http"
	"github.com/uber-go/zap"
	"github.com/shinofara/simple-go-web-app/context"
)

type loggerMiddlewre struct {
	logger zap.Logger
}

func NewLoggerMiddleware() loggerMiddlewre {
	logger := zap.New(
		zap.NewJSONEncoder(JSTTimeFormatter("timestamp")), // drop timestamps in tests
		zap.DebugLevel,
	)
	
	return loggerMiddlewre{
		logger: logger,
	}
}

// loggerMiddlewre.LoggerMiddleware stores Logger to context.
func (ml *loggerMiddlewre) LoggerMiddleware(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) { 
		ctx := context.SetLogger(r.Context(), ml.logger)
		r = r.WithContext(ctx)
		
		ml.logger.Info("Set logger to context.")
			next.ServeHTTP(w, r)		
	}

	return http.HandlerFunc(fn)
}

func JSTTimeFormatter(key string) zap.TimeFormatter {
    return zap.TimeFormatter(func(t time.Time) zap.Field {
        const layout = "2006-01-02 15:04:05"
        jst := time.FixedZone("Asia/Tokyo", 9*60*60)
        return zap.String(key, t.In(jst).Format(layout))
    })
}
