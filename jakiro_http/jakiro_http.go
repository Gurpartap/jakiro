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

package jakiro_http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Gurpartap/jakiro"
	"github.com/gorilla/mux"
)

type HTTPContext struct {
	request *http.Request
	writer  http.ResponseWriter
}

type HTTPErrorResponse struct {
	Errors []jakiro.ContextError `json:"errors"`
}

func NewHTTPContext(writer http.ResponseWriter, request *http.Request) *HTTPContext {
	return &HTTPContext{
		request: request,
		writer:  writer,
	}
}

func (ctx *HTTPContext) Params() map[string]string {
	return mux.Vars(ctx.request)
}

func (ctx *HTTPContext) Body() []byte {
	body, _ := ioutil.ReadAll(ctx.request.Body)
	return body
}

func (ctx *HTTPContext) Write(code int, response []byte) {
	ctx.writer.WriteHeader(code)
	ctx.writer.Write(response)
}

func (ctx *HTTPContext) JSON(code int, object interface{}) {
	var res []byte
	var err error

	if obj, ok := object.(jakiro.JSONEncodable); ok {
		res, err = obj.ToJSON()
	} else {
		res, err = json.MarshalIndent(object, "", "	")
	}

	if err != nil {
		fmt.Println(err.Error())
		ctx.Error(500, err)
		return
	}

	ctx.Write(code, res)
}

func (ctx *HTTPContext) Error(code int, err error) {
	errorResponse := HTTPErrorResponse{}
	errorResponse.Errors = append(errorResponse.Errors, jakiro.ContextError{
		Message: err.Error(),
	})

	res, _ := json.MarshalIndent(errorResponse, "", "	")
	ctx.Write(code, res)
}
