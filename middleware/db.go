package middleware

import (
	"net/http"
	"github.com/shinofara/simple-go-web-app/context"
	"github.com/shinofara/simple-go-web-app/config"		
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	gorp	"gopkg.in/gorp.v1"
	"fmt"
)

// DBMiddleware stores DB connector to context.
func DBMiddleware(dbCfgs *config.DBConfigs) func(http.ResponseWriter, *http.Request, http.HandlerFunc) {

	var dataSourceNames []string
	
	for _, dbCfg := range *dbCfgs {
		dataSourceNames = append(dataSourceNames, fmt.Sprintf("%s:%s@/%s", dbCfg.User, dbCfg.Password, dbCfg.Name))
	}
	
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		ctx := r.Context()

		logger := context.MustGetLogger(ctx)


		db, _ := sql.Open("mysql", dataSourceNames[0])
		dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
		defer dbmap.Db.Close()
		
		ctx = context.SetDB(ctx, dbmap)
		r = r.WithContext(ctx)
		
		logger.Info("Set character string shinofara to context with the name `name`.")

		next(rw, r)
	}
}
