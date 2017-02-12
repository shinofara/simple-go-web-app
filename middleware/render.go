package middleware

import (
	"net/http"
	"github.com/shinofara/simple-go-web-app/context"
	"github.com/shinofara/simple-go-web-app/render"		
)

func RenderMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	re := render.New(rw, r)
	ctx := context.SetRender(r.Context(), re)
	r = r.WithContext(ctx)
	next(rw, r)
}
