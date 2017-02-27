package service

import (
	"fmt"
	"github.com/shinofara/simple-go-web-app/context"
	"github.com/shinofara/simple-go-web-app/repository"
	"github.com/shinofara/simple-go-web-app/entity"
	"github.com/shinofara/simple-go-web-app/transfer"
)

// Register ユーザ登録手続きを行う
func Register(ctx context.Context, user *entity.User) (*entity.User, error) {
	db := context.MustGetDB(ctx, "default")
	logger := context.MustGetLogger(ctx)
	
	//session sample
	sessionStore := context.MustGetSessionStore(ctx)
	login, err := repository.GetLoginSession(sessionStore)
	if err == nil {
		logger.Info(fmt.Sprintf("%+v", login))
	}

	if err != nil {
		logger.Info(err.Error())
	}
	
	_, err = repository.CreateLoginSession(sessionStore)
	if err != nil {
		logger.Info(err.Error())
	}

	//ここでentityと関連付けを行う
	db.AddTableWithName(entity.User{}, "users").SetKeys(true, "ID")

	err = db.CreateTablesIfNotExists()
	if err != nil {
		return	nil, err
	}

	err = repository.CreateUser(db, user)
	if err != nil { 
		return nil, err		
	}

	if err := transfer.SendActivationEmail(ctx); err != nil {
		logger.Error(err.Error())
	}

	return repository.GetUser(db)
}
