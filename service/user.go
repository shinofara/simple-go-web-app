package service

import (
	"github.com/shinofara/simple-go-web-app/context"
	"github.com/shinofara/simple-go-web-app/model/repository"
	"github.com/shinofara/simple-go-web-app/model/entity"
	"github.com/shinofara/simple-go-web-app/transfer"
	"github.com/uber-go/zap"
)

type UserService struct {
	logger zap.Logger
}

func NewUserService(l zap.Logger) *UserService {
	return &UserService{
		logger: l,
	}
}

// Register ユーザ登録手続きを行う
func (us *UserService) Register(ctx context.Context, user *entity.User) (*entity.User, error) {
	db := context.MustGetDB(ctx, "default")

	//ここでentityと関連付けを行う
	db.AddTableWithName(entity.User{}, "users").SetKeys(true, "ID")

	err := db.CreateTablesIfNotExists()
	if err != nil {
		return nil, err
	}

	err = repository.CreateUser(db, user)
	if err != nil {
		return nil, err
	}

	if err := transfer.SendActivationEmail(ctx); err != nil {
		us.logger.Error(err.Error())
	}

	return repository.GetUser(db)
}
