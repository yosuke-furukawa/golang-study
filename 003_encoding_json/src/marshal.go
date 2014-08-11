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
