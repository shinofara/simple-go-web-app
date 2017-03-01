package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"github.com/shinofara/simple-go-web-app/application"
	"github.com/shinofara/simple-go-web-app/config"
	"github.com/shinofara/simple-go-web-app/controller"
	"github.com/shinofara/simple-go-web-app/middleware"
)

// main メイン処理
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
	app.Register("get", "/", controller.Index, []string{"default"})
	app.Register("get", "/example", controller.Example, nil)
	app.Register("get", "/panic", controller.Panic, nil)
	
	// middlewareを登録

	//contextは全体に関わるので一番最初に設定
	app.Router.Use(middleware.ContextMiddleware)
	
	//Loggerは初期化してから追加
	l := middleware.NewLoggerMiddleware()
	app.Router.Use(l.LoggerMiddleware)

	//Loggerは初期化してから追加
	app.Router.Use(middleware.SessionMiddleware("secret"))

	//SampleとRenderは初期化無しで追加
	app.Router.Use(middleware.DBMiddleware(app.Configs, dbCfgs))

	//panic recover
	app.Router.Use(middleware.RecoverMiddleware)

	app.Expand(app.Router)

	log.Fatal(http.ListenAndServeTLS(
		fmt.Sprintf(":%s", cfg.HTTPPort),
		cfg.CertFilePath,
		cfg.KeyFilePath,
		app.Router))
}
