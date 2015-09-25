# Jakiro

Jakiro provides an interface for handling resource actions and JSON encoding/decoding. It also ships with an HTTP and WebSocket implementation of this interface.

In other words, you can use the same handler function to serve http as well as websocket requests. You can also roll your own `jakiro.Context` implementation to serve through another medium.

##### Get

```bash
$ go get github.com/Gurpartap/jakiro
```

##### Usage

See the included example with CRUD examples on the HTTP implementation.

##### License

MIT
