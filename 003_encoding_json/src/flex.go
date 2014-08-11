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
