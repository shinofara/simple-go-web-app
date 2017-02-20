package middleware

import (
	"net/http"
	"github.com/shinofara/simple-go-web-app/application"
	"github.com/shinofara/simple-go-web-app/context"
	"github.com/shinofara/simple-go-web-app/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	gorp	"gopkg.in/gorp.v1"
	"fmt"
)

// DBMiddleware stores DB connector to context.
func DBMiddleware(appCfgs map[string]*application.ApplicationConfig, dbCfgs *config.DBConfigs) func(http.ResponseWriter, *http.Request, http.HandlerFunc) {

	dataSourceNames := convertDBConfigTable(dbCfgs)
	
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		ctx := r.Context()
		logger := context.MustGetLogger(ctx)

		path := r.URL.Path
		//method := r.Method
		appCfg := appCfgs[application.GenerateIndexKey(path)]

		for _, dbCfgName := range appCfg.Databases {
			db, _ := sql.Open("mysql", dataSourceNames[dbCfgName])
			dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
			defer dbmap.Db.Close()

			ctx = context.SetDB(ctx, dbCfgName, dbmap)
		}

		r = r.WithContext(ctx)

		logger.Info("Set character string shinofara to context with the name `name`.")

		next(rw, r)
	}
}

func convertDBConfigTable(dbCfgs *config.DBConfigs) map[string]string {
	dataSourceNames := make(map[string]string)
	
	for key, dbCfg := range *dbCfgs {
		dataSourceNames[key] = fmt.Sprintf("%s:%s@/%s", dbCfg.User, dbCfg.Password, dbCfg.Name)
	}

	return dataSourceNames
}
