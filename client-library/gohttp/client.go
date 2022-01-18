package gohttp

import (
	"net/http"
)

/*
import "github.com/hemuku90/http-client-go/gohttp"
*/

type httpClient struct {
	client  *http.Client
	builder *clientBuilder
}

// Client Interface: Allows consumer to make GET,POST and DELETE API calls
type Client interface {
	Get(url string, headers http.Header) (*Response, error)
	Post(url string, headers http.Header, body interface{}) (*Response, error)
	Delete(url string, headers http.Header) (*Response, error)
}

func (c *httpClient) Get(url string, headers http.Header) (*Response, error) {
	response, err := c.do(http.MethodGet, url, headers, nil)
	return response, err
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*Response, error) {
	response, err := c.do(http.MethodPost, url, headers, body)
	return response, err
}

func (c *httpClient) Delete(url string, headers http.Header) (*Response, error) {
	response, err := c.do(http.MethodGet, url, headers, nil)
	return response, err
}
