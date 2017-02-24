package main

import (
	"flag"
	"fmt"

	"github.com/shinofara/simple-go-web-app/application"
	"github.com/shinofara/simple-go-web-app/config"
	"github.com/shinofara/simple-go-web-app/handler"
	"github.com/shinofara/simple-go-web-app/middleware"
	"github.com/pressly/chi"
	"log"
	"net/http"
)



func main() {
	 var configPath string
	flag.StringVar(&configPath, "conf", "", "path to config yaml path")	
	flag.Parse()
	
	cfg, err := config.Load(configPath)
	if err != nil {
		panic(err)
	}

	dbCfgs, err := config.LoadDBConfig(cfg.DatabaseYmlPath)
	if err != nil {
		panic(err)
	}

	//アプリケーションの管理
	app := application.New()
	app.Register("get", "/", handler.Index, []string{"default"})
	app.Register("get", "/example", handler.Example, nil)	
	
	r := chi.NewRouter()

	// middlewareを登録

	//contextは全体に関わるので一番最初に設定
	r.Use(middleware.ContextMiddleware)
	
	//Loggerは初期化してから追加
	l := middleware.NewLoggerMiddleware()
	r.Use(l.LoggerMiddleware)

	//SampleとRenderは初期化無しで追加
	r.Use(middleware.DBMiddleware(app.ApplicationConfigs, dbCfgs))

	app.Expand(r)

	log.Fatal(http.ListenAndServeTLS(
		fmt.Sprintf(":%s", cfg.HTTPPort),
		cfg.CertFilePath,
		cfg.KeyFilePath,
		r))
}
