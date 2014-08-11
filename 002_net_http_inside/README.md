net/http のより深い所
====================

もう少し深いところについて説明していきます。
一旦一番簡単な実装に戻って、一行ずつ中で何をやってるのか説明していきます。

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
  // 1. HandleFunc
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  // 2. fmt.Fprint
		fmt.Fprint(w, "Hello!")
	})
  // 3. ListenAndServe
	http.ListenAndServe(":4000", nil)
}

```

1. HandleFuncの話
--------------------

HandleFuncの中の実装を追っていくと、以下のHandleという関数に行き着きます。ここでは何をやっているかというと、基本的にはmapに対してkeyにurl patternをvalueにServeHTTPという関数を登録するだけです。

```go
// 与えられたpatternにhandlerを登録します。
// もしhandlerがパターンに既に存在していたら、panicsを呼び出します
(mux *ServeMux) Handle(pattern string, handler Handler) {
        // muxのmapは共有リソースなのでロックを取る
        mux.mu.Lock()
        defer mux.mu.Unlock()

        if pattern == "" {
                panic("http: invalid pattern " + pattern)
        }
        if handler == nil {
                panic("http: nil handler")
        }
        if mux.m[pattern].explicit {
                panic("http: multiple registrations for " + pattern)
        }

        mux.m[pattern] = muxEntry{explicit: true, h: handler, pattern: pattern}

        if pattern[0] != '/' {
                mux.hosts = true
        }

        // 便利なふるまい:
        // もしpatternが/tree/だったら、/treeにリダイレクトする機能がある
        // これはちなみにexplicit registrationによって変更することも可能
        n := len(pattern)
        if n > 0 && pattern[n-1] == '/' && !mux.m[pattern[0:n-1]].explicit {
                path := pattern
                if pattern[0] != '/' {
                        path = pattern[strings.Index(pattern, "/"):]
                }
                // ここが本処理。
                // mux.mというmapにmuxEntry構造体を作って登録する
                mux.m[pattern[0:n-1]] = muxEntry{h: RedirectHandler(path, StatusMovedPermanently), pattern: pattern}
        }
}
```

とまぁこんな感じにHandlerFuncの中ではmuxEntry構造体を作って登録するだけ。

ちなみにHandleFuncは中でDefaultMuxと呼ばれる構造体のServeHTTP関数を呼んでるだけです。HandlerというインタフェースにあるServeHTTP関数を実装しています。このHandleFuncはワザワザ構造体を作らなくてもサーバーを建てられるようにするための糖衣構文と言えるでしょう。

逆に言えばServeHTTP関数を持つ構造体を作って登録すればDefaultMuxのふるまいを変更させることが可能です。

```go
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
```

HandlerFuncのまとめ
-------------------

- やってることはDefaultMuxのServeHTTP関数を呼び出してる
- ServeHTTP関数はResponseWriterとRequestを引数に取り、リクエストが来たら実行される関数
- ということはServeHTTP関数さえ実装すればDefaultMuxの動きを変更できる

2. fmt.Fprintの話
--------------------

fmt.Fprintは第二引数に与えられた文字列を使って第一引数のWriterに対してWrite関数を実行する関数
ということは、Write関数を明示的にこちらから呼べば同じことが実現できる

```go
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
```

fmt.Fprintのまとめ
-------------------

- fmt.Fprintは文字列を受け取ってWriterに渡してるだけ
- Write関数を明示的に呼べば同じことが実現できる


3. ListenAndServeの話
--------------------

httpリクエストは必ずgoroutineが上がるようになっている。

```go
func (srv *Server) Serve(l net.Listener) error {
        defer l.Close()
        var tempDelay time.Duration // how long to sleep on accept failure
        for {
                rw, e := l.Accept()
                if e != nil {
                        if ne, ok := e.(net.Error); ok && ne.Temporary() {
                                if tempDelay == 0 {
                                        tempDelay = 5 * time.Millisecond
                                } else {
                                        tempDelay *= 2
                                }
                                if max := 1 * time.Second; tempDelay > max {
                                        tempDelay = max
                                }
                                srv.logf("http: Accept error: %v; retrying in %v", e, tempDelay)
                                time.Sleep(tempDelay)
                                continue
                        }
                        return e
                }
                tempDelay = 0
                c, err := srv.newConn(rw)
                if err != nil {
                        continue
                }
                c.setState(c.rwc, StateNew) // before Serve can return
                // ここで必ず goroutineが起動する
                go c.serve()
        }
}
```

これにより、リクエストが来る度にgoroutineが起動するようになっており、mainルーチンをブロックしない作りになっている。


