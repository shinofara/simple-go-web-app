package middleware

import (
	"net/http"
	"github.com/shinofara/simple-go-web-app/session"	
	"github.com/shinofara/simple-go-web-app/context"
)

func SessionMiddleware(secret string) func (next http.Handler) http.Handler {
	return func (next http.Handler) http.Handler  {
		fn := func(w http.ResponseWriter, r *http.Request) {
			sess := session.New(w, r, secret)
			ctx := context.SetSession(r.Context(), sess)
			r = r.WithContext(ctx)
			
			next.ServeHTTP(w, r)		
		}

		return http.HandlerFunc(fn)
	}
}
