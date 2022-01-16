package gohttp

import (
	"net/http"
)

/*
import "github.com/hemuku90/http-client-go/gohttp"
func main(){
	gohttp.Client{}
}
*/

func NewClient() HttpClient {
	client := &httpClient{}
	return client
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
