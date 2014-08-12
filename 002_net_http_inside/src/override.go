package main

import (
	"fmt"
	"net/http"
)

type greeter struct{}

// ServeHTTPを実装する
// Golangでは、ServeHTTP関数を実装するだけでHandlerインタフェースを実装したことになる
func (g *greeter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	// ListenAndServeの第二引数に構造体を作って渡す
	http.ListenAndServe(":4000", &greeter{})
}
