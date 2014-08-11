package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:name`
	Age  int    `json:age`
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
