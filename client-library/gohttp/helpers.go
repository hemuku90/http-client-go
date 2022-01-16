package gohttp

import (
	"net/http"
	"time"
)

func (c *httpClient) SetHeaders(headers http.Header) {
	c.headers = headers
}

func (c *httpClient) SetConnectionTimeout(timeout time.Duration) {
	c.connectionTimeout = timeout
}
func (c *httpClient) SetRequestTimeout(timeout time.Duration) {
	c.responseTimeout = timeout
}
func (c *httpClient) SetMaxIdleConenctionPerHost(maxIdleConenctionPerHost int) {
	c.maxIdleConnections = maxIdleConenctionPerHost
}
