// Package render 表示に関するパッケージ
package render

import (
	"github.com/unrolled/render"
	"github.com/shinofara/simple-go-web-app/http/context"
	"net/http"
	"go.uber.org/zap"
)

// Render 表示に必要な情報を管理
type Render struct {
	w http.ResponseWriter
	r *http.Request
	render *render.Render
	logger *zap.SugaredLogger
}

// New creates a Render
func New(w http.ResponseWriter, r *http.Request) *Render {
	re := render.New(render.Options{
		DisableHTTPErrorRendering: true,
		Directory: "resource/template",
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

// HTML 指定されたテンプレートにデータバインドを行って出力
func (r *Render) HTML(name string, data map[string]string) {
	err := r.render.HTML(r.w, http.StatusOK, name, data)
	if err != nil{
		http.Redirect(r.w, r.r, "/notfound", http.StatusFound)
	}
}
