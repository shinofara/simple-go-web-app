package middleware

import (
	"net/http"
	"context"
	"time"
)

func ContextMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	//1.7からはbackgroudだけではなく、net/httpにもcontextが追加されたので、それを利用する
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 10 * time.Second)
	defer cancel()
	r = r.WithContext(ctx)
  next(rw, r)		
}
