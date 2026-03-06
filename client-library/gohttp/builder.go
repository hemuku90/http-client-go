package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	headers            http.Header
	maxIdleConnections int
	disableTimeout     bool
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
}

type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	DisableTimeouts(disableTimeouts bool) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetRequestTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConenctionPerHost(maxIdleConenctionPerHost int) ClientBuilder
	Build() Client
}

func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}
	return builder
}

func (c *clientBuilder) Build() Client {
	client := httpClient{
		builder: c,
	}
	return &client
}

/* Setting Client Configuration with clientBuilder*/
func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}
func (c *clientBuilder) SetRequestTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout
	return c
}
func (c *clientBuilder) SetMaxIdleConenctionPerHost(maxIdleConnectionPerHost int) ClientBuilder {
	c.maxIdleConnections = maxIdleConnectionPerHost
	return c
}

func (c *clientBuilder) DisableTimeouts(disableTimeouts bool) ClientBuilder {
	c.disableTimeout = disableTimeouts
	return c
}
