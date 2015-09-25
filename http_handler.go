package bebe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPHandler struct {
	request *http.Request
	writer  http.ResponseWriter
}

func NewHTTPHandler(writer http.ResponseWriter, request *http.Request) *HTTPHandler {
	return &HTTPHandler{
		request: request,
		writer:  writer,
	}
}

func (h *HTTPHandler) Params() map[string]string {
	fmt.Println("h.request", h.request)
	return mux.Vars(h.request)
}

func (h *HTTPHandler) Body() []byte {
	body, _ := ioutil.ReadAll(h.request.Body)
	return body
}

func (h *HTTPHandler) Success(code int, message string) {
	h.Write(code, []byte(message))
}

func (h *HTTPHandler) Error(code int, err error) {
	h.Write(code, []byte(err.Error()))
}

func (h *HTTPHandler) JSON(code int, object interface{}) {
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

func (h *HTTPHandler) Write(code int, response []byte) {
	h.writer.WriteHeader(code)
	h.writer.Write(response)
}
