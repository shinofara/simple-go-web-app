package controller

import (
	"github.com/shinofara/simple-go-web-app/render"
	"net/http"
)

// Example Get:/exampleの処理を定義
func Example(rw http.ResponseWriter, r *http.Request) {
	re := render.New(rw, r)
	re.HTML("sample", map[string]string{"name": "example"})
	return
}
