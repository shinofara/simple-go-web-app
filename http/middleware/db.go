package middleware

import (
	"net/http"
	"github.com/shinofara/simple-go-web-app/router"
	"github.com/shinofara/simple-go-web-app/http/context"
	"github.com/shinofara/simple-go-web-app/config"
	"github.com/shinofara/simple-go-web-app/model/entity"
	"database/sql"

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
	gorp	"gopkg.in/gorp.v1"
	"fmt"
)

// DBMiddleware stores DB connector to context.
func DBMiddleware(appCfgs router.Configs, dbCfgs *config.DBConfigs) func(next http.Handler) http.Handler {
	dataSourceNames := convertDBConfigTable(dbCfgs)

	return func(next http.Handler) http.Handler {

		fn := dbMiddleware(next, appCfgs, dataSourceNames)
		return http.HandlerFunc(fn)
	}
}

// dbMiddleware http.Handler
func dbMiddleware(next http.Handler, appCfgs router.Configs, dataSourceNames map[string]string) func(rw http.ResponseWriter, r *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		path := r.URL.Path
		//method := r.Method

		appCfg := appCfgs.GetPathConfig(path)

		if appCfg != nil {
			for _, dbCfgName := range appCfg.Databases {
				db, _ := sql.Open("mysql", dataSourceNames[dbCfgName])
				dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
				defer dbmap.Db.Close()

				associateTable(dbmap)

				ctx = context.SetDB(ctx, dbCfgName, dbmap)
			}
		}

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
	}
}

// convertDBConfigTable yamlの構成を使いやすい形に変換
func convertDBConfigTable(dbCfgs *config.DBConfigs) map[string]string {
	dataSourceNames := make(map[string]string)
	
	for key, dbCfg := range *dbCfgs {
		dataSourceNames[key] = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Name)
	}

	return dataSourceNames
}

func associateTable(db *gorp.DbMap) {
	db.AddTableWithName(entity.User{}, "users").SetKeys(true, "ID")
}
