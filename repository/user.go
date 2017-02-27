package repository

import (
	"encoding/base64"
	"github.com/shinofara/simple-go-web-app/entity"
	gorp "gopkg.in/gorp.v1"
)

type secretString string

// Encrypt 元となる文字列を暗号化する
func (s secretString) Encrypt() string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// CreateUser ユーザを登録
func CreateUser(db *gorp.DbMap, name string, pass secretString) error {
	inv2 := &entity.User{
		Name: name,
		Pass: pass.Encrypt(),
	}

	// Insert your rows
	return db.Insert(inv2)
}

// GetUser ユーザを取得
func GetUser(db *gorp.DbMap) (*entity.User, error) {
	user, err := db.Get(entity.User{}, 1)
	if err != nil {
		return nil, err
	}
	return user.(*entity.User), nil
}
