package middleware

import (
	"net/http"
	"context"
	"time"
)

func ContextMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		//1.7からはbackgroudだけではなく、net/httpにもcontextが追加されたので、それを利用する
		ctx := r.Context()
		ctx, cancel := context.WithTimeout(ctx, 10 * time.Second)
		defer cancel()
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
