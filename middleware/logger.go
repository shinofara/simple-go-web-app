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

func (ml *loggerMiddlewre) LoggerMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx := context.SetLogger(r.Context(), ml.logger)
	r = r.WithContext(ctx)
	
	ml.logger.Info("Set logger to context.")
	next(rw, r)
}

func JSTTimeFormatter(key string) zap.TimeFormatter {
    return zap.TimeFormatter(func(t time.Time) zap.Field {
        const layout = "2006-01-02 15:04:05"
        jst := time.FixedZone("Asia/Tokyo", 9*60*60)
        return zap.String(key, t.In(jst).Format(layout))
    })
}
