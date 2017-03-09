// Package router アプリケーション内のroute毎に必要な設定などを行う
package router

//applicationの全体設定に必要な定義を行う
//route設定と、route毎のDB設定管理を担う

import (
	"fmt"
	"regexp"
	"net/http"
	"github.com/pressly/chi"	
)

// Application routeとそれに対応する設定等を管理
type Router struct {
	Mux *chi.Mux
	tree map[string]*Method
	Configs map[string]*Config
}

// Method http method毎のpathを保持
type Method struct {
	path map[string]http.HandlerFunc
}

// Configs mapを型として定義
type Configs map[string]*Config

// Config 各pathが必要とする設定を保持
// 現時点では使用するDB名のみ
type Config struct {
	Key string
	Databases []string
}

// New creates a Application.
func New() *Router {
	return &Router{
		Mux: chi.NewRouter(),
		tree: make(map[string]*Method),
		Configs: make(map[string]*Config),
	}
}

// Register method毎のpathに対応したhandlerを登録
func (a *Router) Register(method, path string, handler http.HandlerFunc, databases []string) {
	if a.tree[method] == nil {
		a.tree[method] = new(Method)
	}

	if a.tree[method].path == nil {
		a.tree[method].path = make(map[string]http.HandlerFunc)
	}

	a.tree[method].path[path] = handler

	key := generateIndexKey(path)
	a.Configs[key] = &Config{Key: key, Databases: databases}
}



func (r *Router) Middleware(f func(http.Handler) http.Handler) {
	r.Mux.Use(f)
}

func (r *Router) Router() *chi.Mux {
	r.expand()
	return r.Mux
}

// Expand muxに保持しているpathを展開
func (a *Router) expand() {
	for method, paths := range a.tree {
		for path, handler := range paths.path {
			if method == "get" {
				a.Mux.Get(path, handler)
			}
		}
	}
}

// GetPathConfig pathに対応したconfigを取得
func (a *Configs) GetPathConfig(path string) *Config {
	return (*a)[generateIndexKey(path)]
}

// GenerateIndexKey pathに対応したkeyを生成
func generateIndexKey(path string) string {
	re := regexp.MustCompile("(.*)/$")
  path = re.ReplaceAllString(path, "$1")
	return fmt.Sprintf("%s%s", "get", path)
}
