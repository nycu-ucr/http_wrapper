package http_wrapper

import (
	"github.com/nycu-ucr/gonet/http"
)

type Response struct {
	Header http.Header
	Status int
	Body   interface{}
}

func NewResponse(code int, h http.Header, body interface{}) (ret *Response) {
	ret = &Response{}
	ret.Status = code
	ret.Header = h
	ret.Body = body
	return
}
