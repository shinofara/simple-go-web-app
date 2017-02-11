package middleware

import (
	"net/http"
	"github.com/unrolled/render"
	"context"
)

func RenderMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	re := NewRender(rw, r)
	ctx := r.Context()
	ctx = context.WithValue(ctx, "render", re)	
	r = r.WithContext(ctx)
  next(rw, r)	
}

type Render struct {
	w http.ResponseWriter
	r *http.Request	
	render *render.Render

}

func NewRender(w http.ResponseWriter, r *http.Request) *Render {
	re := render.New(render.Options{
		DisableHTTPErrorRendering: true,
	})
	
	return &Render{
		w: w,
		r: r,
		render: re,
	}
}

func (r *Render) HTML(name string, data map[string]string) {

	err := r.render.HTML(r.w, http.StatusOK, name, data)
	if err != nil{
		http.Redirect(r.w, r.r, "/notfound", http.StatusFound)
	}
}
