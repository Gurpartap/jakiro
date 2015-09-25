package bebe

import (
	"encoding/json"
	"fmt"

	"golang.org/x/net/websocket"
)

type WebsocketRequest struct {
	Params map[string]string `json:"params"`
	Body   []byte            `json:"body"`
}

type WebsocketResponse struct {
	Status int    `json:"status"`
	Body   []byte `json:"body"`
}

type WebsocketHandler struct {
	request WebsocketRequest
	ws      *websocket.Conn
}

func NewWebsocketHandler(request WebsocketRequest, ws *websocket.Conn) *WebsocketHandler {
	return &WebsocketHandler{
		request: request,
		ws:      ws,
	}
}

func (h *WebsocketHandler) Params() map[string]string {
	return h.request.Params
}

func (h *WebsocketHandler) Body() []byte {
	return []byte("")
}

func (h *WebsocketHandler) Success(code int, message string) {
	h.Write(code, []byte(message))
}

func (h *WebsocketHandler) Error(code int, err error) {
	h.Write(code, []byte(err.Error()))
}

func (h *WebsocketHandler) JSON(code int, object interface{}) {
	var res []byte
	var err error

	if obj, ok := object.(JSONEncodable); ok {
		res, err = obj.ToJSON()
	} else {
		res, err = json.MarshalIndent(object, "", "	")
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	h.Write(code, res)
}

func (h *WebsocketHandler) Write(code int, response []byte) {
	wsResponse := WebsocketResponse{Status: code, Body: response}
	res, err := json.MarshalIndent(wsResponse, "", "	")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = websocket.Message.Send(h.ws, res)

	if err != nil {
		fmt.Println(err.Error())
	}
}
