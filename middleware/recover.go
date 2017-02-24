package middleware

import (
	"github.com/shinofara/simple-go-web-app/context"
	"runtime/debug"	
	"net/http"
	"fmt"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	fn := func(rw http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {

				ctx := r.Context()
				logger := context.MustGetLogger(ctx)
				logger.Error(fmt.Sprintf("%s %s", rvr, debug.Stack()))

				http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(rw, r)
	}

	return http.HandlerFunc(fn)
}
