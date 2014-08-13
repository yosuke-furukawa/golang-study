package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
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
