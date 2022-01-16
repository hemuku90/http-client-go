package gohttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	defaultmaxIdleConnections = 5
	defaultresponseTimeout    = 5 * time.Second
	defaultconnectionTimeout  = 2 * time.Second
)

func (c *httpClient) do(httpMethod string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	allHeaders := c.GetRequestHeaders(headers)
	requestBody, err := c.GetRequestBody(allHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal the request Body")
	}
	request, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("unable to create request to %s", url)
	}
	request.Header = allHeaders //Set custom and common headers for the request
	client := c.getHTTPClient()
	response, err := client.Do(request)
	return response, err
}

//getHTTPClient: returns *http.Client
func (c *httpClient) getHTTPClient() *http.Client {
	if c.client != nil {
		return c.client
	}
	client := &http.Client{
		Timeout: c.getconnectionTimeout() + c.getresponseTimeout(),
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   c.getmaxIdleConnections(),
			ResponseHeaderTimeout: c.getresponseTimeout(), //Response timeout after request is send
			DialContext: (&net.Dialer{
				Timeout: c.getconnectionTimeout(), // Socket connection timeout
			}).DialContext,
		},
	}
	c.client = client
	return c.client
}

//GetRequestBody: returns []byte,error
func (c *httpClient) GetRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	switch strings.ToLower(contentType) { // Add support for different content-types
	case "application/json":
		return json.Marshal(body)
	default:
		return json.Marshal(body)
	}
}

func (c *httpClient) getmaxIdleConnections() int {
	if c.maxIdleConnections > 0 {
		return c.maxIdleConnections
	}
	//Disable Timeouts
	if c.disableTimeout {
		return 0
	}
	return defaultmaxIdleConnections
}

func (c *httpClient) getresponseTimeout() time.Duration {
	if c.responseTimeout > 0 {
		return c.responseTimeout
	}
	//Disable Timeouts
	if c.disableTimeout {
		return 0
	}
	return defaultresponseTimeout
}

func (c *httpClient) getconnectionTimeout() time.Duration {
	if c.connectionTimeout > 0 {
		return c.connectionTimeout
	}
	//Disable Timeouts
	if c.disableTimeout {
		return 0
	}
	return defaultconnectionTimeout
}

func (c *httpClient) DisableTimeouts(disableTimeouts bool) {
	c.disableTimeout = disableTimeouts
}

//getRequestHeaders returns custom headers from do() and common headers for httpClient instance
func (c *httpClient) GetRequestHeaders(requestHeaders http.Header) http.Header {
	finalHeaders := make(http.Header)
	//Default headers for the request
	for header, value := range c.headers {
		if len(value) > 0 {
			finalHeaders.Set(header, value[0])
		}
	}
	//Custom headers for the request
	for header, value := range requestHeaders {
		if len(value) > 0 {
			finalHeaders.Set(header, value[0])
		}
	}
	return finalHeaders
}
