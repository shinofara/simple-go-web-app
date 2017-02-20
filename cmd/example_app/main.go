package main

import (
	"flag"
	"fmt"

	"github.com/shinofara/simple-go-web-app/config"	
	"github.com/shinofara/simple-go-web-app/router"	
	"github.com/shinofara/simple-go-web-app/middleware"
	"github.com/urfave/negroni"
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

	dbCfgs, err := config.LoadDBConfig("./database.yml")
	if err != nil {
		panic(err)
	}
	
	n := negroni.New()

	// middlewareを登録

	//contextは全体に関わるので一番最初に設定
	n.Use(negroni.HandlerFunc(middleware.ContextMiddleware))
	
	//Loggerは初期化してから追加
	l := middleware.NewLoggerMiddleware()
	n.Use(negroni.HandlerFunc(l.LoggerMiddleware))

	//SampleとRenderは初期化無しで追加
	n.Use(negroni.HandlerFunc(middleware.DBMiddleware(dbCfgs)))

	// http handlerを登録
	r := router.New()
	n.UseHandler(r.Router)

	log.Fatal(http.ListenAndServeTLS(
		fmt.Sprintf(":%s", cfg.HTTPPort),
		cfg.CertFilePath,
		cfg.KeyFilePath,
		n))
}
