package service

import (
	"net/http"	
	"github.com/shinofara/simple-go-web-app/entity"
	"github.com/shinofara/simple-go-web-app/repository"
	"github.com/shinofara/simple-go-web-app/context"		
)

func CreateNewUser(r *http.Request, name string) (*entity.User, error) {
	dbmap := context.MustGetDB(r)

	//ここでentityと関連付けを行う
	dbmap.AddTableWithName(entity.User{}, "users").SetKeys(true, "ID")

	err := dbmap.CreateTablesIfNotExists()
	if err != nil {
		return	nil, err
	}

	err = repository.CreateUser(dbmap, name)
	if err != nil {
		return nil, err		
	}

	return repository.GetUser(dbmap)
}
