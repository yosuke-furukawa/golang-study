package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	// HandleFuncを使ってリクエストとレスポンスを処理します
	// ここでは、"/" 以下全てのリクエストをHelloで返します。
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		path = strings.Replace(path, "/", " ", -1)
		fmt.Fprintf(w, "Hello!%s", path)
	})
	// 4000 portで起動
	http.ListenAndServe(":4000", nil)
}
