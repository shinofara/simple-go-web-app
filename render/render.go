package render

import (
	"github.com/unrolled/render"
	"github.com/shinofara/simple-go-web-app/context"
	"net/http"
	"github.com/uber-go/zap"		
)

type Render struct {
	w http.ResponseWriter
	r *http.Request	
	render *render.Render
	logger zap.Logger
}

func New(w http.ResponseWriter, r *http.Request) *Render {
	re := render.New(render.Options{
		DisableHTTPErrorRendering: true,
		Directory: "template",
		Charset: "UTF-8",
		HTMLContentType: "text/html",
	})

	logger := context.MustGetLogger(r.Context())
	
	return &Render{
		w: w,
		r: r,
		render: re,
		logger: logger,
	}
}

func (r *Render) HTML(name string, data map[string]string) {
	err := r.render.HTML(r.w, http.StatusOK, name, data)
	if err != nil{
		http.Redirect(r.w, r.r, "/notfound", http.StatusFound)
	}
}
