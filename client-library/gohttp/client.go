package gohttp

import (
	"net"
	"net/http"
	"time"
)

/*
import "github.com/hemuku90/http-client-go/gohttp"
func main(){
	gohttp.Client{}
}
*/

func NewClient() HttpClient {
	client := &httpClient{
		client: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   5,
				ResponseHeaderTimeout: 5 * time.Second, //Response timeout after request is send
				DialContext: (&net.Dialer{
					Timeout: 5 * time.Second, // Socket connection timeout
				}).DialContext,
			},
		},
	}
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
