package handlers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/shinofara/simple-go-web-app/render"
	"net/http"
)

func Example(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	re := render.New(rw, r)
	re.HTML("sample", map[string]string{"name": "example"})
	return
}
