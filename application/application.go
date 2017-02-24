package application

//applicationの全体設定に必要な定義を行う
//route設定と、route毎のDB設定管理を担う

import (
	"fmt"
	"regexp"
	"net/http"
	"github.com/pressly/chi"	
)

type Application struct {
	Router map[string]*Method
	ApplicationConfigs map[string]*ApplicationConfig
}

type Method struct {
	path map[string]http.HandlerFunc
}

type ApplicationConfig struct {
	Key string
	Databases []string
}

func New() *Application {
	return &Application{
		Router: make(map[string]*Method),
		ApplicationConfigs: make(map[string]*ApplicationConfig),
	}
}

func (a *Application) Register(method, path string, handler http.HandlerFunc, databases []string) {
	if a.Router[method] == nil {
		a.Router[method] = new(Method)
	}

	if a.Router[method].path == nil {
		a.Router[method].path = make(map[string]http.HandlerFunc)
	}

	a.Router[method].path[path] = handler


	key := GenerateIndexKey(path)
	a.ApplicationConfigs[key] = &ApplicationConfig{Key: key, Databases: databases}
}

func (a *Application) Expand(mx *chi.Mux) {
	for method, paths := range a.Router {
		for path, handler := range paths.path {
			if method == "get" {
				mx.Get(path, handler)
			}
		}
	}
}

func GenerateIndexKey(path string) string {
	re := regexp.MustCompile("(.*)/$")
  path = re.ReplaceAllString(path, "$1")
	return fmt.Sprintf("%s%s", "get", path)
}
