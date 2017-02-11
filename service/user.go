package service

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shinofara/simple-go-web-app/entity"
	"github.com/shinofara/simple-go-web-app/repository"	
	"gopkg.in/gorp.v1"	
)

func CreateNewUser(name string) (*entity.User, error) {
	db, err := sql.Open("mysql", "root:test@/test")
	if err != nil {
		return	nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	defer dbmap.Db.Close()
	
	//ここでentityと関連付けを行う
	dbmap.AddTableWithName(entity.User{}, "users").SetKeys(true, "ID")

	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		return	nil, err
	}

	err = repository.CreateUser(dbmap, name)
	if err != nil {
		return nil, err		
	}

	return repository.GetUser(dbmap)
}
