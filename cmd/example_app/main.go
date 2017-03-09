package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"github.com/shinofara/simple-go-web-app/config"
	"github.com/shinofara/simple-go-web-app/controller"
	"github.com/shinofara/simple-go-web-app/http/router"
	"github.com/shinofara/simple-go-web-app/http/middleware"
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
	r:= router.New()
	r.Register("get", "/", controller.Example, []string{"default"})
	r.Register("get", "/panic", controller.Panic, nil)

	// middlewareを登録

	//contextは全体に関わるので一番最初に設定
	r.Middleware(middleware.ContextMiddleware)

	//Loggerは初期化してから追加
	l := middleware.NewLoggerMiddleware()
	r.Middleware(l.LoggerMiddleware)

	//Loggerは初期化してから追加
	r.Middleware(middleware.SessionMiddleware("salt"))

	//SampleとRenderは初期化無しで追加
	r.Middleware(middleware.DBMiddleware(r.Configs, dbCfgs))

	//panic recover
	r.Middleware(middleware.RecoverMiddleware)

	log.Fatal(http.ListenAndServeTLS(
		fmt.Sprintf(":%s", cfg.HTTPPort),
		cfg.CertFilePath,
		cfg.KeyFilePath,
		r.Router()))
}
