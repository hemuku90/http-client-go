package gohttp

import (
	"net/http"
)

type httpClient struct {
	client  *http.Client
	headers http.Header
	// maxIdleconnections int
	// connectionTimeout  time.Duration
	// requestTimeout     time.Duration
}

type HttpClient interface {
	SetHeaders(headers http.Header)
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}
