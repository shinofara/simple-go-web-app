package handler

import (
	"github.com/shinofara/simple-go-web-app/render"
	"github.com/shinofara/simple-go-web-app/service"
	"github.com/shinofara/simple-go-web-app/context"	
	"net/http"
)

func Index(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := context.MustGetLogger(ctx)
	us := service.NewUser(ctx)
	user, err := us.Register("test")

	re := render.New(rw, r)
	if err != nil {
		logger.Error(err.Error())
		re.HTML("sample", map[string]string{"name": err.Error()})
		return
	}
	re.HTML("sample", map[string]string{"name": user.Name})
	return
}
