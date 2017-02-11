package middleware

import (
	"net/http"
	"github.com/uber-go/zap"
	"github.com/shinofara/simple-go-web-app/context"	
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	gorp	"gopkg.in/gorp.v1"
)

func DBMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx := r.Context()

	logger := ctx.Value("logger").(zap.Logger)

	db, _ := sql.Open("mysql", "root:test@/test")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	defer dbmap.Db.Close()

	ctx = context.SetDB(ctx, dbmap)
	r = r.WithContext(ctx)

	logger.Info("Set character string shinofara to context with the name `name`.")
  next(rw, r)
}
