package middleware

import (
	"net/http"
	"github.com/uber-go/zap"
	_ "github.com/go-sql-driver/mysql"
	"context"
)

func SampleMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx := r.Context()
	logger := ctx.Value("logger").(zap.Logger)

	ctx = context.WithValue(ctx, "name", "shinofara")
	r = r.WithContext(ctx)	

	logger.Info("Set character string shinofara to context with the name `name`.")
  next(rw, r)	
}
