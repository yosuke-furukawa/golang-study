net/http
===================

Goのhttpネットワークに関する機能を提供する標準パッケージです。
基本的にhttpリクエストを発行する場合やhttpサーバを起動するときにはこのパッケージを利用します。

説明を読むよりも手を動かしていきましょう。

httpサーバーを建てる
-------------------

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello!")
	})
	http.ListenAndServe(":4000", nil)
}
```

これが一番簡単なHTTPサーバーです。
もう少し説明を追記します。

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
  // HandleFuncを使ってリクエストとレスポンスを処理します
  // ここでは、"/" 以下全てのリクエストをHelloで返します。
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello!")
	})
  // 4000 portで起動
	http.ListenAndServe(":4000", nil)
}
```

じゃあこれをechoサーバにするため、urlのパスを取ってHelloにくっつけるようにしてみましょう。

```
# イメージ
$ curl http://localhost:4000/world
# Hello! world
```

では作っていきましょう

```go

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
    // pathを取るにはr.URL.Pathで受け取る
		path := r.URL.Path
    // stringsパッケージのReplaceで/を空白に変えておく
		path = strings.Replace(path, "/", " ", -1)
    // fmt.Fprintfに変更する
		fmt.Fprintf(w, "Hello!%s", path)
	})
	// 4000 portで起動
	http.ListenAndServe(":4000", nil)
}
```


こんな感じに成ります。

実際にブラウザもしくはcurlでアクセスしてみてください。

```
$ curl http://localhost:4000/world
# Hello! world
$ curl http://localhost:4000/world/yosuke
# Hello! world yosuke
```

ここから更にエンドポイントを増やしていきましょう

```go
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
	// 新しくエンドポイントを追加する
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
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
```

こうすると、QueryParameterを解すようになります。

```
$ curl http://localhost:4000/api?name=yosuke&age=31
# Hello! yosuke, your age is 31
```

URLは以下の様な構造をしています。

```go
type URL struct {
        Scheme   string
        Opaque   string
        User     *Userinfo
        Host     string
        Path     string
        RawQuery string
        Fragment string
}
```

それぞれ以下の様なURL構造にマッチしています。

```go
scheme://[userinfo@]host/path[?query][#fragment]
```

詳しくはURLの公式ドキュメントを見るとよいでしょう。

[http://golang.org/pkg/net/url/](http://golang.org/pkg/net/url/)

まとめ
-----------------

- httpサーバを建てる
- HandleFuncを使ってエンドポイントを増やす
- クエリーパラメータやurlのパスを得るにはURL structを使う
