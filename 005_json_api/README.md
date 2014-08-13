json apiを作ろう
================

さて、ここまでの話を基に総合してjson の APIを作ってみましょう。


仕様：
/apiから始まるurlの時にquery paramから来た値をファイルに保存します。
/からその値をjsonにして取り出してみましょう。
jsonは属性が決められたもので良いです。
もちろん時間があればフレキシブルなJSONを保持してそれを取り出すようにしても構いません。

例：


```
$ curl http://localhost:4000/api?name=hoge&age=20
{ "name" : "hoge", "age" : "20" }
$ curl http://localhost:4000/
{ "name" : "hoge", "age" : "20" }
```
