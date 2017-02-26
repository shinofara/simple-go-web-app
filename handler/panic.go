package handler

import (
	"net/http"
)

// Panic Get:/panicの処理を定義
func Panic(_ http.ResponseWriter, _ *http.Request) {
	panic("recover example")
}
