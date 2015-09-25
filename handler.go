package bebe

type Handler interface {
	Params() map[string]string
	Body() []byte

	Success(code int, message string)
	Error(code int, err error)
	JSON(code int, object interface{})
	Write(code int, response []byte)
}
