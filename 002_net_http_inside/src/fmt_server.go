package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// WriterのWrite関数を読んであげるだけで良い。
		// Write関数はbyteの配列を受け付けるのでstringからキャストする必要がある
		w.Write([]byte("Hello!!!"))
	})
	http.ListenAndServe(":4000", nil)
}
