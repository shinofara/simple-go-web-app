package application

//applicationの全体設定に必要な定義を行う
//route設定と、route毎のDB設定管理を担う

import (
	"fmt"
	"regexp"
	"github.com/julienschmidt/httprouter"
)

type Application struct {
	Router *httprouter.Router
	ApplicationConfigs map[string]*ApplicationConfig
}

type ApplicationConfig struct {
	Key string
	Databases []string
}

func New() *Application {
	router := httprouter.New()

	return &Application{
		Router: router,
		ApplicationConfigs: make(map[string]*ApplicationConfig),
	}
}

func (a *Application) Register(path string, handler httprouter.Handle, databases []string) {
	a.Router.GET(path, handler)

	key := GenerateIndexKey(path)
	a.ApplicationConfigs[key] = &ApplicationConfig{Key: key, Databases: databases}
}

func GenerateIndexKey(path string) string {
	re := regexp.MustCompile("(.*)/$")
  path = re.ReplaceAllString(path, "$1")
	return fmt.Sprintf("%s%s", "get", path)
}
