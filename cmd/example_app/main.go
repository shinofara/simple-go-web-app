package main

import (
	"flag"
	"log"
	"net"
	"net/http/fcgi"
	"github.com/pressly/chi"
	"github.com/shinofara/simple-go-web-app/config"
	"github.com/shinofara/simple-go-web-app/controller"
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
	r := chi.NewRouter()

	//contextは全体に関わるので一番最初に設定
	r.Use(middleware.ContextMiddleware)

	//Loggerは初期化してから追加
	l := middleware.NewLoggerMiddleware()
	r.Use(l.LoggerMiddleware)

	//Loggerは初期化してから追加
	r.Use(middleware.SessionMiddleware("salt"))

	//SampleとRenderは初期化無しで追加
	r.Use(middleware.DBMiddleware(dbCfgs))

	//panic recover
	r.Use(middleware.RecoverMiddleware)

	r.Get("/", controller.Example)
	r.Get("/panic", controller.Panic)	

	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	log.Fatal(fcgi.Serve(listen, r))
}
