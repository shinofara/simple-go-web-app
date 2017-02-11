package service

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shinofara/simple-go-web-app/entity"
	"gopkg.in/gorp.v1"	
)

func CreateNewUser(name string) (*entity.User, error) {
	db, err := sql.Open("mysql", "root:test@/test")
	if err != nil {
		return	nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	//ここでentityと関連付けを行う
	dbmap.AddTableWithName(entity.User{}, "users").SetKeys(true, "ID")

	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		return	nil, err
	}
	
	defer dbmap.Db.Close()

	inv2 := &entity.User{Name: name}

	// Insert your rows
	err = dbmap.Insert(inv2)
	if err != nil {
		return nil, err
	}

	obj, err := dbmap.Get(entity.User{}, 8)
	if err != nil {
		return nil, err
	}

	return obj.(*entity.User), nil
}
