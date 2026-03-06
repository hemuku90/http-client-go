package gohttp

import (
	"net/http"
)

type Response struct {
	headers    http.Header
	statusCode int
	body       []byte
}

func (r *Response) GetHeaders() http.Header {
	return r.headers
}

func (r *Response) GetStatusCode() int {
	return r.statusCode
}

func (r *Response) String() string {
	return string(r.body)
}
func (r *Response) Bytes() []byte {
	return r.body
}
