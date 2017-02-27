package repository

import (
	"github.com/shinofara/simple-go-web-app/entity"
	gorp "gopkg.in/gorp.v1"
)

// CreateUser ユーザを登録
func CreateUser(db *gorp.DbMap, user *entity.User) error {
	// Insert your rows
	return db.Insert(user)
}

// GetUser ユーザを取得
func GetUser(db *gorp.DbMap) (*entity.User, error) {
	user, err := db.Get(entity.User{}, 1)
	if err != nil {
		return nil, err
	}
	return user.(*entity.User), nil
}
