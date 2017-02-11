package main

import (
	"flag"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/shinofara/simple-go-web-app/handlers"
	"github.com/shinofara/simple-go-web-app/middleware"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)
var (
	CertFilePath string
	KeyFilePath string
	HTTPPort string
)

func init() {
	flag.StringVar(&CertFilePath, "ssl-cert", "", "path to cert file")
	flag.StringVar(&KeyFilePath, "ssl-key", "", "path to key file")
	flag.StringVar(&HTTPPort, "http-port", "8080", "numbuer of port")	

	flag.Parse()
}

func main() {
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
		fmt.Sprintf(":%s", HTTPPort),
		CertFilePath,
		KeyFilePath,
		n))
}
