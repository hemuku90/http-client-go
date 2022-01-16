package gohttp

import (
	"net/http"
	"time"
)

type httpClient struct {
	client             *http.Client
	headers            http.Header
	maxIdleConnections int
	disableTimeout     bool
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
}

type HttpClient interface {
	SetHeaders(headers http.Header)
	DisableTimeouts(disableTimeouts bool)
	SetConnectionTimeout(timeout time.Duration)
	SetRequestTimeout(timeout time.Duration)
	SetMaxIdleConenctionPerHost(maxIdleConenctionPerHost int)
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}
