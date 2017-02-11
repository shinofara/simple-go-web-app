package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shinofara/simple-go-web-app/entity"
	gorp "gopkg.in/gorp.v1"	
)

func CreateUser(db *gorp.DbMap, name string) error {
	inv2 := &entity.User{Name: name}

	// Insert your rows
	return db.Insert(inv2)
}

func GetUser(db *gorp.DbMap) (*entity.User, error) {
	user, err := db.Get(entity.User{}, 8)
	if err != nil {
		return nil, err
	}
	return user.(*entity.User), nil	
}
