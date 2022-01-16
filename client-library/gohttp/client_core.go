package gohttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *httpClient) GetRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)
	default:
		return json.Marshal(body)
	}
}

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
	//Set custom and common headers for the request
	request.Header = allHeaders
	response, err := c.client.Do(request)
	return response, err
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
