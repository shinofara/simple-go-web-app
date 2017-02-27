package handler

import (
	"github.com/shinofara/simple-go-web-app/render"
	"github.com/shinofara/simple-go-web-app/service"
	"github.com/shinofara/simple-go-web-app/context"
	"github.com/shinofara/simple-go-web-app/entity"	
	"github.com/gorilla/schema"
	"net/http"
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

	user := entity.NewUser(u.Name, "password")
	
	user, err = service.Register(ctx, user)

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
