package middleware

import (
	"net/http"
	"github.com/uber-go/zap"
	"github.com/nbio/httpcontext"	
)

func SampleMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	logger := httpcontext.Get(r, "logger").(zap.Logger)
	httpcontext.Set(r, "name", "shinofara")

	logger.Info("Set character string shinofara to context with the name `name`.")
  next(rw, r)	
}
