package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	ioutil.WriteFile("./hoge.txt", []byte("Hello"), 0644)
	fmt.Println("Wrote.")
	t, err := ioutil.ReadFile("./hoge.txt")
	if err != nil {
		fmt.Errorf("%s\n", err)
	}
	fmt.Printf("READ : %s\n", t)
}
