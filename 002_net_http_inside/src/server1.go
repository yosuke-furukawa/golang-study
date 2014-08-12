package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		path = strings.Replace(path, "/", " ", -1)
		fmt.Fprintf(w, "Hello!%s", path)
	})
	// 新しくエンドポイントを追加する
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// QueryParameterを取る
		// Mapの形式で帰ってくる
		queryParam := r.URL.Query()
		// Mapにはkey-valueで入っていて、valueにはGetでアクセスする
		name := queryParam.Get("name")
		age := queryParam.Get("age")
		fmt.Fprintf(w, "Hello! %s, your age is %s", name, age)
	})
	// 4000 portで起動
	http.ListenAndServe(":4000", nil)
}
