package service

import (
	"github.com/shinofara/simple-go-web-app/context"
	"github.com/shinofara/simple-go-web-app/model/entity"
	"github.com/shinofara/simple-go-web-app/model/transfer"
	"github.com/shinofara/simple-go-web-app/model/repository"
)

// UserService ユーザたいする振る舞い
type UserService struct {}

// NewUserService creates a UserService
func NewUserService() *UserService {
	return &UserService{}
}

// Register ユーザ登録手続きを行う
func (us *UserService) Register(ctx context.Context, user *entity.User) (*entity.User, error) {
	db := context.MustGetDB(ctx, "default")
	logger := context.MustGetLogger(ctx)	

	err := db.CreateTablesIfNotExists()
	if err != nil {
		return nil, err
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
