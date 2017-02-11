package middleware

import (
	"net/http"
	"github.com/uber-go/zap"
	"github.com/nbio/httpcontext"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	gorp	"gopkg.in/gorp.v1"	
)

func SampleMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	logger := httpcontext.Get(r, "logger").(zap.Logger)
	httpcontext.Set(r, "name", "shinofara")

	logger.Info("Set character string shinofara to context with the name `name`.")
  next(rw, r)	
}

func DBMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	logger := httpcontext.Get(r, "logger").(zap.Logger)

	db, _ := sql.Open("mysql", "root:test@/test")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	defer dbmap.Db.Close()

	httpcontext.Set(r, "DB", dbmap)

	logger.Info("Set character string shinofara to context with the name `name`.")
  next(rw, r)
}
