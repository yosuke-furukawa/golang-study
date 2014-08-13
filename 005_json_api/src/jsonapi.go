// 解答
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u User) String() string {
	return fmt.Sprintf(`{ "name": %s, "age" : %d }`, u.Name, u.Age)
}

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		queryParam := r.URL.Query()
		name := queryParam.Get("name")
		age, _ := strconv.Atoi(queryParam.Get("age"))
		user := &User{Name: name, Age: age}
		userJson, _ := json.Marshal(user)
		ioutil.WriteFile("./user.json", userJson, 0644)
		w.Header().Set("Content-type", "application/json")
		fmt.Fprint(w, string(userJson))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		userJson, _ := ioutil.ReadFile("./user.json")
		w.Header().Set("Content-type", "application/json")
		fmt.Fprint(w, string(userJson))
	})
	// 4000 portで起動
	http.ListenAndServe(":4000", nil)
}
