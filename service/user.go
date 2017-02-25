package service

import (
	"github.com/shinofara/simple-go-web-app/context"
	"github.com/shinofara/simple-go-web-app/repository"
	"github.com/shinofara/simple-go-web-app/entity"
	"github.com/shinofara/simple-go-web-app/transfer"
)

type UserService struct {
	ctx context.Context
}

func NewUser(ctx context.Context) *UserService {
	return &UserService{ctx}
}

func (u *UserService) Register(name string) (*entity.User, error) {
	db := context.MustGetDB(u.ctx, "default")
	logger := context.MustGetLogger(u.ctx)

	//ここでentityと関連付けを行う
	db.AddTableWithName(entity.User{}, "users").SetKeys(true, "ID")

	err := db.CreateTablesIfNotExists()
	if err != nil {
		return	nil, err
	}

	err = repository.CreateUser(db, name)
	if err != nil {
		return nil, err		
	}

	if err := transfer.SendActivationEmail(u.ctx); err != nil {
		logger.Error(err.Error())
	}

	return repository.GetUser(db)
}
