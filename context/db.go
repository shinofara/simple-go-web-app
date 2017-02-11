package context

import (
	"net/http"
	"github.com/nbio/httpcontext"
	gorp	"gopkg.in/gorp.v1"
	"fmt"
)

func SetDB(r *http.Request, db *gorp.DbMap) {
	httpcontext.Set(r, "DB", db)	
}

func GetDB(r *http.Request) (*gorp.DbMap, error) {
	db, ok :=  httpcontext.Get(r, "DB").(*gorp.DbMap)
	if ok {
		return db, nil
	}
	
	return nil, fmt.Errorf("Failed to get DB from context")
}

func MustGetDB(r *http.Request) *gorp.DbMap {
	db, err := GetDB(r)
	if err != nil {
		panic(err)
	}
	return db
}
