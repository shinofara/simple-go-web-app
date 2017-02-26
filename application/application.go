// Package application アプリケーション内のroute毎に必要な設定などを行う
package application

//applicationの全体設定に必要な定義を行う
//route設定と、route毎のDB設定管理を担う

import (
	"fmt"
	"regexp"
	"net/http"
	"github.com/pressly/chi"	
)

// Application routeとそれに対応する設定等を管理
type Application struct {
	Router *chi.Mux
	tree map[string]*Method
	ApplicationConfigs map[string]*ApplicationConfig
}

// Method http method毎のpathを保持
type Method struct {
	path map[string]http.HandlerFunc
}

// ApplicationConfigs mapを型として定義
type ApplicationConfigs map[string]*ApplicationConfig

// ApplicationConfig 各pathが必要とする設定を保持
// 現時点では使用するDB名のみ
type ApplicationConfig struct {
	Key string
	Databases []string
}

// New creates a Application.
func New() *Application {
	return &Application{
		Router: chi.NewRouter(),
		tree: make(map[string]*Method),
		ApplicationConfigs: make(map[string]*ApplicationConfig),
	}
}

// Register method毎のpathに対応したhandlerを登録
func (a *Application) Register(method, path string, handler http.HandlerFunc, databases []string) {
	if a.tree[method] == nil {
		a.tree[method] = new(Method)
	}

	if a.tree[method].path == nil {
		a.tree[method].path = make(map[string]http.HandlerFunc)
	}

	a.tree[method].path[path] = handler

	key := generateIndexKey(path)
	a.ApplicationConfigs[key] = &ApplicationConfig{Key: key, Databases: databases}
}

// Expand muxに保持しているpathを展開
func (a *Application) Expand(mx *chi.Mux) {
	for method, paths := range a.tree {
		for path, handler := range paths.path {
			if method == "get" {
				mx.Get(path, handler)
			}
		}
	}
}

// GetPathConfig pathに対応したconfigを取得
func (a *ApplicationConfigs) GetPathConfig(path string) *ApplicationConfig {
	return (*a)[generateIndexKey(path)]
}

// GenerateIndexKey pathに対応したkeyを生成
func generateIndexKey(path string) string {
	re := regexp.MustCompile("(.*)/$")
  path = re.ReplaceAllString(path, "$1")
	return fmt.Sprintf("%s%s", "get", path)
}
