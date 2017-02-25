package handler

import (
	"github.com/shinofara/simple-go-web-app/render"
	"github.com/shinofara/simple-go-web-app/service"
	"github.com/shinofara/simple-go-web-app/context"
	"github.com/shinofara/simple-go-web-app/session"
	"github.com/gorilla/schema"
	"net/http"
)

type User struct {
	Name  string
}

func Index(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// Handle error
	}
	var u User

	var decoder = schema.NewDecoder()
	err = decoder.Decode(&u, r.Form)
	if err != nil {
		// Handle error
	}

	//session
	s := session.New(rw, r, "secret")
	s.SetLoginData()
	
	ctx := r.Context()
	logger := context.MustGetLogger(ctx)
	us := service.NewUser(ctx)
	user, err := us.Register(u.Name)

	re := render.New(rw, r)
	if err != nil {
		logger.Error(err.Error())
		re.HTML("sample", map[string]string{"name": err.Error()})
		return
	}
	user.Name = u.Name
	re.HTML("sample", map[string]string{"name": user.Name})
	return
}
