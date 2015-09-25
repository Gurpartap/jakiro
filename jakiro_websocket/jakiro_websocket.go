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

package jakiro_websocket

import (
	"encoding/json"
	"fmt"

	"github.com/Gurpartap/jakiro"
	"github.com/gorilla/websocket"
)

type WebSocketRequest struct {
	Params map[string]string `json:"params"`
	Body   []byte            `json:"body"`
}

type WebSocketResponse struct {
	Status int    `json:"status"`
	Body   []byte `json:"body"`
}

type WebSocketErrorResponse struct {
	Status int                   `json:"status"`
	Errors []jakiro.ContextError `json:"errors"`
}

type WebSocketContext struct {
	request WebSocketRequest
	ws      *websocket.Conn
}

func NewWebSocketContext(request WebSocketRequest, ws *websocket.Conn) *WebSocketContext {
	return &WebSocketContext{
		request: request,
		ws:      ws,
	}
}

func (ctx *WebSocketContext) Params() map[string]string {
	return ctx.request.Params
}

func (ctx *WebSocketContext) Body() []byte {
	return ctx.request.Body
}

func (ctx *WebSocketContext) Write(code int, response []byte) {
	ctx.ws.WriteJSON(WebSocketResponse{Status: code, Body: response})
}

func (ctx *WebSocketContext) JSON(code int, object interface{}) {
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

	ctx.ws.WriteJSON(WebSocketResponse{Status: code, Body: res})
}

func (ctx *WebSocketContext) Error(code int, err error) {
	errorResponse := WebSocketErrorResponse{Status: code}
	errorResponse.Errors = append(errorResponse.Errors, jakiro.ContextError{
		Message: err.Error(),
	})

	ctx.ws.WriteJSON(errorResponse)
}
