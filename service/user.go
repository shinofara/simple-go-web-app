package service

import (
	"net/http"	
	"github.com/nbio/httpcontext"
	"github.com/shinofara/simple-go-web-app/entity"
	"github.com/shinofara/simple-go-web-app/repository"	
	gorp	"gopkg.in/gorp.v1"	
)

func CreateNewUser(r *http.Request, name string) (*entity.User, error) {
	dbmap := httpcontext.Get(r, "DB").(*gorp.DbMap)	

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
