package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/nbio/httpcontext"
	"github.com/urfave/negroni"
	
	"net/http"
	"log"

	"simplego/middleware"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	name := httpcontext.GetString(r, "name")
	re := httpcontext.Get(r, "render").(*middleware.Render)

	re.HTML("sample", map[string]string{"name": name})
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

	log.Fatal(http.ListenAndServe(":8080", n))
}
