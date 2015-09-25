# jakiro
[![GoDoc](https://godoc.org/github.com/Gurpartap/jakiro?status.svg)](https://godoc.org/github.com/Gurpartap/jakiro)

![](http://i1.2pcdn.com/node14/201401/25/article_sub_img/a0doyp0jezdubwfo/2.jpg)

> Image from [dotafire.com](http://www.dotafire.com/dota-2/guide/sandos-guide-to-jakiro-9689). Go there to learn about Jakiro, the dual breath dragon!

`jakiro` provides a lightweight interface for handling resource actions and JSON encoding/decoding. It also ships with an HTTP and WebSocket implementation of this interface.

In other words, you can use the same handler function to serve http as well as websocket requests. You can also roll your own `jakiro.Context` implementation to serve through another medium.

```go
type Context interface {
	Params() map[string]string
	Body() []byte

	Write(code int, response []byte)
	JSON(code int, object interface{})
	Error(code int, err error)
}
```

#### Implementations

`jakiro` ships with two implementations of `jakiro.Context`: [`HTTPContext`](https://github.com/Gurpartap/jakiro/blob/master/http_context.go) and [`WebSocketContext`](https://github.com/Gurpartap/jakiro/blob/master/websocket_context.go).

`HTTPContext`'s implementation of `Params()` assumes that you're using [gorilla/mux](https://github.com/gorilla/mux).

Similarly, `WebSocketContext` assumes [gorilla/websocket](https://github.com/gorilla/websocket).

Fork away for your choice of socket, RPC or HTTP request multiplexer! Implementations are portable anyway!

#### Get

```bash
go get github.com/Gurpartap/jakiro
```

#### Usage

```go
package main

import (
	"net/http"

	"github.com/Gurpartap/jakiro"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

var DBUsersTable = make(map[int64]*User, 0)

func withJakiro(handlerFunc func(jakiro.Context)) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		handlerFunc(jakiro.NewHTTPContext(rw, req))
	}
}

func main() {
	r := mux.NewRouter()
	n := negroni.Classic()

	r.Methods("GET").Path("/api/users").HandlerFunc(withJakiro(IndexUserHandler))
	r.Methods("POST").Path("/api/users").HandlerFunc(withJakiro(CreateUserHandler))
	r.Methods("GET").Path("/api/users/{id:[0-9]+}").HandlerFunc(withJakiro(ReadUserHandler))
	r.Methods("DELETE").Path("/api/users/{id:[0-9]+}").HandlerFunc(withJakiro(DestroyUserHandler))

	n.UseHandler(r)
	n.Run(":3000")
}
```

See **rest of the [example](https://github.com/Gurpartap/jakiro/tree/master/example) code** with CRUD handlers and JSON interfaces usage.

#### Questions? Bugs?

[Create a new issue](https://github.com/Gurpartap/jakiro/issues/new)

#### License

MIT
