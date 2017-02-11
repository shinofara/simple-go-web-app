package handlers

import (

	"github.com/julienschmidt/httprouter"
	"github.com/shinofara/simple-go-web-app/service"

	"net/http"
	"github.com/shinofara/simple-go-web-app/middleware"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()	
	re := ctx.Value("render").(*middleware.Render)
	name := ctx.Value("name").(string)		

	u, err := service.CreateNewUser(r, name)
  if err != nil {
		re.HTML("sample", map[string]string{"name": err.Error()})		
	}

	re.HTML("sample", map[string]string{"name": u.Name})
}
