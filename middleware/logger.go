package middleware

import (
	"time"
	"net/http"
	"github.com/uber-go/zap"
	"github.com/nbio/httpcontext"	
)

func LoggerMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	logger := zap.New(
    zap.NewJSONEncoder(JSTTimeFormatter("timestamp")), // drop timestamps in tests
	)	
	httpcontext.Set(r, "logger", logger)
	logger.Info("Set logger to context.")
  next(rw, r)	
}

func JSTTimeFormatter(key string) zap.TimeFormatter {
    return zap.TimeFormatter(func(t time.Time) zap.Field {
        const layout = "2006-01-02 15:04:05"
        jst := time.FixedZone("Asia/Tokyo", 9*60*60)
        return zap.String(key, t.In(jst).Format(layout))
    })
}
