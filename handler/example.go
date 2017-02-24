package handler

import (
	"github.com/shinofara/simple-go-web-app/render"
	"net/http"
)

func Example(rw http.ResponseWriter, r *http.Request) {
	re := render.New(rw, r)
	re.HTML("sample", map[string]string{"name": "example"})
	return
}
