package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/nbio/httpcontext"
	"github.com/urfave/negroni"
	
	"net/http"
	"log"
	"flag"
	"fmt"

	"github.com/shinofara/simple-go-web-app/middleware"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	name := httpcontext.GetString(r, "name")
	re := httpcontext.Get(r, "render").(*middleware.Render)

	re.HTML("sample", map[string]string{"name": name})
}

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
	n.Use(negroni.HandlerFunc(middleware.LoggerMiddleware))
	n.Use(negroni.HandlerFunc(middleware.SampleMiddleware))
	n.Use(negroni.HandlerFunc(middleware.RenderMiddleware))	


  // http handlerを登録
	router := httprouter.New()
	router.GET("/", Index)
  n.UseHandler(router)

	log.Fatal(http.ListenAndServeTLS(
		fmt.Sprintf(":%s", HTTPPort),
		CertFilePath,
		KeyFilePath,
		n))
}
