package httpcli

import "net/http"

type HttpReq struct {
	Method string
	Header http.Header
}

func NewHttpReq(method string) *HttpReq {
	return &HttpReq{
		Method: method,
	}
}
