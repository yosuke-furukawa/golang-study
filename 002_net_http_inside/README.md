net/http のより深い所
====================

httpリクエストは必ずgoroutineが上がるようになっている。

```go
func (srv *Server) Serve(l net.Listener) error {
        defer l.Close()
        var tempDelay time.Duration // how long to sleep on accept failure
        for {
                rw, e := l.Accept()
                if e != nil {
                        if ne, ok := e.(net.Error); ok && ne.Temporary() {
                                if tempDelay == 0 {
                                        tempDelay = 5 * time.Millisecond
                                } else {
                                        tempDelay *= 2
                                }
                                if max := 1 * time.Second; tempDelay > max {
                                        tempDelay = max
                                }
                                srv.logf("http: Accept error: %v; retrying in %v", e, tempDelay)
                                time.Sleep(tempDelay)
                                continue
                        }
                        return e
                }
                tempDelay = 0
                c, err := srv.newConn(rw)
                if err != nil {
                        continue
                }
                c.setState(c.rwc, StateNew) // before Serve can return
                // ここで必ず goroutineが起動する
                go c.serve()
        }
}
```

これにより、リクエストが来る度にgoroutineが起動するようになっており、mainルーチンをブロックしない作りになっている。
