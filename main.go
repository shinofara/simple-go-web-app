package main

import (
	"flag"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/shinofara/simple-go-web-app/config"	
	"github.com/shinofara/simple-go-web-app/handlers"
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
	
	n := negroni.New()

	// middlewareを登録

	//contextは全体に関わるので一番最初に設定
	n.Use(negroni.HandlerFunc(middleware.ContextMiddleware))
	
	//Loggerは初期化してから追加
	l := middleware.NewLoggerMiddleware()
	n.Use(negroni.HandlerFunc(l.LoggerMiddleware))

	//SampleとRenderは初期化無しで追加
	n.Use(negroni.HandlerFunc(middleware.DBMiddleware))	
	n.Use(negroni.HandlerFunc(middleware.SampleMiddleware))
	n.Use(negroni.HandlerFunc(middleware.RenderMiddleware))	

	// http handlerを登録
	router := httprouter.New()
	router.GET("/", handlers.Index)
	n.UseHandler(router)

	log.Fatal(http.ListenAndServeTLS(
		fmt.Sprintf(":%s", cfg.HTTPPort),
		cfg.CertFilePath,
		cfg.KeyFilePath,
		n))
}
