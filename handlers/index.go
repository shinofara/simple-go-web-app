package handlers

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/shinofara/simple-go-web-app/entity"
	//"github.com/shinofara/simple-go-web-app/middleware"
	"gopkg.in/gorp.v1"
	"net/http"
	"github.com/nbio/httpcontext"
	"github.com/shinofara/simple-go-web-app/middleware"
	"fmt"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	re := httpcontext.Get(r, "render").(*middleware.Render)
	
	db, err := sql.Open("mysql", "root:test@/test")
	if err != nil {
		re.HTML("sample", map[string]string{"name": err.Error()})
		return		
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	//ここでentityと関連付けを行う
	dbmap.AddTableWithName(entity.User{}, "users").SetKeys(true, "ID")

	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		re.HTML("sample", map[string]string{"name": "err0" + err.Error()})
		return		
	}
	
	defer dbmap.Db.Close()

	inv2 := &entity.User{ID: 1, Name: "test"}

	// Insert your rows
	err = dbmap.Insert(inv2)
	if err != nil {
		re.HTML("sample", map[string]string{"name": "err1" + err.Error()})
		return
	}

	obj, err := dbmap.Get(entity.User{}, 1)
	if err != nil {
		re.HTML("sample", map[string]string{"name": "err2"+  err.Error()})
		return
	}

	inv := obj.(*entity.User)

	name := httpcontext.GetString(r, "name")


	re.HTML("sample", map[string]string{"name": fmt.Sprintf("%s%s", inv.Name, name)})
}
