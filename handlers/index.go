package handlers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/shinofara/simple-go-web-app/render"
	"github.com/shinofara/simple-go-web-app/service"
	"net/http"
)

func Index(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u, err := service.CreateNewUser(r.Context(), "test")

	re := render.New(rw, r)
	if err != nil {
		re.HTML("sample", map[string]string{"name": err.Error()})
		return
	}
	re.HTML("sample", map[string]string{"name": u.Name})
	return
}
