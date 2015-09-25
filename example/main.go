// Copyright (c) 2015 Gurpartap Singh
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"net/http"

	"github.com/Gurpartap/jakiro"
	"github.com/Gurpartap/jakiro/jakiro_http"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

var DBUsersTable = make(map[int64]*User, 0)

func withJakiro(handlerFunc func(jakiro.Context)) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		handlerFunc(jakiro_http.NewHTTPContext(rw, req))
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
