package controller

import (
	"fmt"
	"github.com/shinofara/simple-go-web-app/app/render"
	"github.com/shinofara/simple-go-web-app/context"
	"github.com/shinofara/simple-go-web-app/http/request"
	"github.com/shinofara/simple-go-web-app/model/entity"
	"github.com/shinofara/simple-go-web-app/model/service"
	"github.com/shinofara/simple-go-web-app/session"
	"net/http"
)

// User ユーザ情報
type User struct {
	Name  string
}

// Index Get:/の処理を定義
func Example(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := context.MustGetLogger(ctx)
	
	var u User
	if err := request.Decode(r, &u); err != nil {
		logger.Error(err.Error())
	}

	logger.Infow("Failed to fetch URL.",
		"key", "value",
	)
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
	
	userService := service.NewUserService()
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
