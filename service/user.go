package service

import (
	"github.com/shinofara/simple-go-web-app/context"
	"github.com/shinofara/simple-go-web-app/repository"
	"github.com/shinofara/simple-go-web-app/entity"
)

type UserService struct {
	ctx context.Context
}

func NewUser(ctx context.Context) *UserService {
	return &UserService{ctx}
}

func (u *UserService) Register(name string) (*entity.User, error) {
	db := context.MustGetDB(u.ctx, "default")

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

	return repository.GetUser(db)
}
