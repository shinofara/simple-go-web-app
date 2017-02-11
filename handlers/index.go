package handlers

import (

	"github.com/julienschmidt/httprouter"
	"github.com/shinofara/simple-go-web-app/service"

	"net/http"
	"github.com/nbio/httpcontext"
	"github.com/shinofara/simple-go-web-app/middleware"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	re := httpcontext.Get(r, "render").(*middleware.Render)

	name := httpcontext.GetString(r, "name")

	u, err := service.CreateNewUser(name)
  if err != nil {
		re.HTML("sample", map[string]string{"name": err.Error()})		
	}

	re.HTML("sample", map[string]string{"name": u.Name})
}
