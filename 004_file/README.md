io package
=================

さて、ioパッケージを使うとファイルへのRead/Writeを行うことができ、簡易的な永続化を行うことが可能です。

ioの使い方を説明していきます。

fileに書く
-----------------

fileに書くにはio/ioutilパッケージを使うと簡単に実施することが可能です。

```go
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	ioutil.WriteFile("./hoge.txt", []byte("Hello"), 0644)
	fmt.Println("DONE.")
}

```

書いたら読んでみたいですよね。

fileを読む
------------------

fileを読むのもioutil.ReadFileを使うと簡単です。

```go
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
```


他にもBufferを使ったやり方があるのですが、現時点では割愛します。
そのうち追記します。
