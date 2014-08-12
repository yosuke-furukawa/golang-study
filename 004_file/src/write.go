package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	ioutil.WriteFile("./hoge.txt", []byte("Hello"), 0644)
	fmt.Println("DONE.")
}
