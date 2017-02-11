package render

import (
	"github.com/unrolled/render"
	"net/http"
)

type Render struct {
	w http.ResponseWriter
	r *http.Request	
	render *render.Render
}

func New(w http.ResponseWriter, r *http.Request) *Render {
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
