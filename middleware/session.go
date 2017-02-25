package middleware

import (
	"net/http"
	"github.com/shinofara/simple-go-web-app/session"	
	"github.com/shinofara/simple-go-web-app/context"
)

func SessionMiddleware(secret string) func (next http.Handler) http.Handler {
	return func (next http.Handler) http.Handler  {
		fn := func(w http.ResponseWriter, r *http.Request) {
			store := session.NewSessionStore(w, r, secret)
			ctx := context.SetSessionStore(r.Context(), store)
			r = r.WithContext(ctx)
			
			next.ServeHTTP(w, r)		
		}

		return http.HandlerFunc(fn)
	}
}
