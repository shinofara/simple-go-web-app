package controller

import (
	"github.com/shinofara/simple-go-web-app/context"
	"github.com/shinofara/simple-go-web-app/app/render"
	"github.com/shinofara/simple-go-web-app/model/service"
	"github.com/shinofara/simple-go-web-app/model/entity"
	"github.com/shinofara/simple-go-web-app/session"
	"github.com/gorilla/schema"
	"net/http"
	"fmt"
)

// User ユーザ情報
type User struct {
	Name  string
}

// Index Get:/の処理を定義
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

	ctx := r.Context()
	logger := context.MustGetLogger(ctx)
	sessionStore := context.MustGetSessionStore(ctx)

	//session
	login, err := session.GetLoginSession(sessionStore)
	if err != nil {
		logger.Info(err.Error())
	}
	logger.Info(fmt.Sprintf("%+v", login))

	_, err = session.CreateLoginSession(sessionStore)
	if err != nil {
		logger.Info(err.Error())
	}

	user := entity.NewUser(u.Name, "password")
	
	userService := service.NewUserService(logger)
	user, err = userService.Register(ctx, user)

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
