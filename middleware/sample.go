package middleware

import (
	"net/http"
	"github.com/shinofara/simple-go-web-app/context"
)

func SampleMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx := r.Context()
	logger := context.MustGetLogger(ctx)

	logger.Info("Set character string shinofara to context with the name `name`.")
	next(rw, r)
}
