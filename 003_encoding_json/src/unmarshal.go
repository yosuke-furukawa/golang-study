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
