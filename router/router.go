package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/shinofara/simple-go-web-app/handlers"	
)

type Router struct {
	Router *httprouter.Router
}

type ResourceCfg struct {
	rw string
}

func New() *Router {
	r := &Router{
		Router: httprouter.New(),
	}

	r.Router.GET("/", handlers.Index)
	
	return r

}
