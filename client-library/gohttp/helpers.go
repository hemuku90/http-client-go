package gohttp

import "net/http"

func (c *httpClient) SetHeaders(headers http.Header) {
	c.headers = headers
}
