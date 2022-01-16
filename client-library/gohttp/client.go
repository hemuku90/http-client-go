package gohttp

import "net/http"

/*
import "github.com/hemuku90/http-client-go/gohttp"
func main(){
	gohttp.Client{}
}
*/

type httpClient struct {
	client  *http.Client
	headers http.Header
}

func NewClient() HttpClient {
	client := &httpClient{
		client: &http.Client{},
	}
	return client
}

type HttpClient interface {
	SetHeaders(headers http.Header)
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	response, err := c.do(http.MethodGet, url, headers, nil)
	return response, err
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	response, err := c.do(http.MethodPost, url, headers, body)
	return response, err
}

func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	response, err := c.do(http.MethodGet, url, headers, nil)
	return response, err
}

func (c *httpClient) SetHeaders(headers http.Header) {
	c.headers = headers
}
