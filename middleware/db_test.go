package middleware

import (
	"reflect"
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/pressly/chi"		
	"github.com/shinofara/simple-go-web-app/config"
	"github.com/shinofara/simple-go-web-app/application"
	"github.com/shinofara/simple-go-web-app/context"
)

func TestDBMiddleware(t *testing.T) {
	appCfg := map[string]*application.ApplicationConfig{
		"get": &application.ApplicationConfig{
			Key: "get",
			Databases: []string{"default", "read"},
		},
	}

	dbCfgs := &config.DBConfigs{
		"default": &config.DBConfig{
			Name: "test",
			User: "test",
			Password: "test",
			Host: "localhost",
			Port: 3306,
		},
		"read": &config.DBConfig{
			Name: "read_test",
			User: "read_test",
			Password: "read_test",
			Host: "read_host",
			Port: 3306,
		},		
	}

	//仮想のリクエストを生成
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	//仮想のリクエストハンドラを生成
	testHandler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if _, err := context.GetDB(r.Context(), "default"); err !=nil {
			t.Fatalf("Must exists default db connection in Context. %s", err.Error())			
		}

		if _, err := context.GetDB(r.Context(), "read"); err !=nil {
			t.Fatalf("Must exists read db connection in Context. %s", err.Error())
		}
	})

	//テスト用httpサーバを立ち上げ、テストリクエストを実行
	r := chi.NewRouter()
	r.Use(ContextMiddleware)
	l := NewLoggerMiddleware()	
	r.Use(l.LoggerMiddleware)
	r.Use(DBMiddleware(appCfg, dbCfgs))
	r.Get("/", testHandler)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, request)
}

func TestConvertDBConfigTable(t *testing.T) {
	dbCfgs := &config.DBConfigs{
		"default": &config.DBConfig{
			Name: "test",
			User: "test",
			Password: "test",
			Host: "localhost",
			Port: 3306,
		},
	}

	dbTables := convertDBConfigTable(dbCfgs)

	expected := map[string]string{"default": "test:test@tcp(localhost:3306)/test"}

	if !reflect.DeepEqual(dbTables, expected) {
		t.Errorf("Must be equal, \ne is %+v \na is %+v", expected, dbTables)
	}
	
}
