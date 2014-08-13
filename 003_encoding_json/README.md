encoding/json
===================

httpサーバは建てられるようになったので、JSON APIを作るために必要なencoding/jsonの使い方を学んでいきましょう。

encoding/jsonを使うとjsonを生成することやjson形式の文字列からgolangの構造体を作ることができるようになります。

基本的には2つの関数さえ覚えておけば使えると思います。

Marshal関数
------------------

JSONのbyte列を生成する関数です。使ってみましょう。

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	// ここでUserを作る
	m := User{"yosuke furukawa", 31}
	// Marshal関数でbyteの文字列を作る
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	// byteの文字列を出力する
	fmt.Printf("%s\n", string(b))
}
```

このファイルを go run で実行させると以下の様な結果が出ると思います。

```
$ go run marshal.go
{"Name":"yosuke furukawa","Age":31}
```

Unmarshal関数
------------------

Unmarshalはbyte列を構造体に変換する関数です。

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	var user User
	// Unmarshalを使うと構造体にマッピングできる
	err := json.Unmarshal([]byte(`{"Name":"yosuke furukawa", "Age":31}`), &user)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	//構造体の値を表示する
	fmt.Printf("%s\n", user.Name)
	fmt.Printf("%d\n", user.Age)
}
```

```
$ go run unmarshal.go
yosuke furukawa
31
```

さて、このままではjsonの文字列が構造体の変数名とマッピングされているので、外部からも見える変数として使いたいけどjsonのkeyは小文字にしたい場合や、jsonのkeyと変数名は違うものにしたい場合に困ります。

こういう時はstructにアノテーションを書きます。

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
  // json:から始まるannotationを書く
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var user User
	// Unmarshalを使うと構造体にマッピングできる
	// アノテーションを使うことでname, ageでもName, Ageにマッピングできる
	err := json.Unmarshal([]byte(`{"name":"yosuke furukawa", "age":31}`), &user)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	//構造体の値を表示する
	fmt.Printf("%s\n", user.Name)
	fmt.Printf("%d\n", user.Age)
}
```

tips structにはStringerインタフェースを実装させると便利。
-----------------------

JavaでいうtoStringのようにgolangにもString関数を実装させると文字列として評価する時に使えます。早速使ってみましょう。


```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// String関数を実装する
func (u User) String() string {
	return fmt.Sprintf("my name is %s, my age is %d", u.Name, u.Age)
}

func main() {
	var user User
	// Unmarshalを使うと構造体にマッピングできる
	// アノテーションを使うことでname, ageでもName, Ageにマッピングできる
	err := json.Unmarshal([]byte(`{"name":"yosuke furukawa", "age":31}`), &user)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	//構造体の値を表示する
	//my name is yosuke furukawa, my age is 31
	fmt.Printf("%s\n", user)
}
```

jsonの型がフレキシブルでよくわからん時
-------------------------

これもありますよね。jsonで型作ろうとするけど、フレキシブルでスキーマがよく分からない時。こういう時は`interface{}`を使います。

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// userの型がわからないので、ここではinterface型にする
	var user interface{}
	err := json.Unmarshal([]byte(`{"name":"yosuke furukawa", "age":31}`), &user)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	// とりあえず値を表示する
	// map[name:yosuke furukawa age:31]
	fmt.Printf("%v\n", user)
	// interfaceからmapにして、値を取り出す
	m := user.(map[string]interface{})
	fmt.Printf("%s\n", m["name"])
	fmt.Printf("%v\n", m["age"])
}
```

参考資料
------------------

- [http://golang.org/doc/articles/json_and_go.html](http://golang.org/doc/articles/json_and_go.html)

