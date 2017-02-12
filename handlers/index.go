package handlers

import (

	"github.com/julienschmidt/httprouter"
	"github.com/shinofara/simple-go-web-app/service"

	"net/http"
	"github.com/shinofara/simple-go-web-app/context"		
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	re := context.MustGetRender(r.Context())

	u, err := service.CreateNewUser(r, "test")
	if err != nil {
		re.HTML("sample", map[string]string{"name": err.Error()})		
	}
	re.HTML("sample", map[string]string{"name": u.Name})
}
