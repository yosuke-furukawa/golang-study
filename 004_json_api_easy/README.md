json apiを作ろう
================

さて、ここまでの話を基に総合してjson の APIを作ってみましょう。
ひとまず永続化は無視してメモリ中に保存するだけで考えてみます。


仕様：
/apiから始まるurlの時にquery paramから来た値を保存します。
/からその値をjsonにして取り出してみましょう。
何のパラメータが来ても 保存できるようにinterfaceを使いましょう。
値はすべて文字列で良いですが、もしも時間があれば型を見て変換してください。

例：


```
$ curl http://localhost:4000/api?name=hoge&age=20
{ "name" : "hoge", "age" : "20" }
$ curl http://localhost:4000/
{ "name" : "hoge", "age" : "20" } #前と同じリクエストが帰る
$ curl http://localhost:4000/api?name=hoge&age=20&hoge=fuga
{ "name" : "hoge", "age" : "20", "hoge" : "fuga" }
$ curl http://localhost:4000/
{ "name" : "hoge", "age" : "20", "hoge" : "fuga" } #上書きできる
```
