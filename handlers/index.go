package handlers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/nbio/httpcontext"
	"github.com/shinofara/simple-go-web-app/middleware"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	name := httpcontext.GetString(r, "name")
	re := httpcontext.Get(r, "render").(*middleware.Render)

	re.HTML("sample", map[string]string{"name": name})
}
