package gohttp

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	//Initialisation
	client := httpClient{}
	headers := make(http.Header)
	headers.Set("Content-Type", "application/json")
	headers.Set("User-Agent", "http-client")
	client.builder.headers = headers
	//Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("Custom-Header", "Custom-HTTP-Client-Header")
	actualHeaders := client.GetRequestHeaders(requestHeaders)
	actualHeadersCount := len(actualHeaders)
	expectedHeadersCount := 3
	//Result
	if actualHeadersCount != expectedHeadersCount {
		t.Errorf("Expected %v headers but got %v headers", actualHeadersCount, expectedHeadersCount)
	}
	if actualHeaders.Get("Custom-Header") != "Custom-HTTP-Client-Header" {
		t.Errorf("Expected Custom-Header value to be: Custom-HTTP-Client-Header ")
	}
}

func TestGetRequestBody(t *testing.T) {
	client := httpClient{}
	//SubTest for Testing Nil body
	t.Run("TestGetRequestBodyNil", func(t *testing.T) {
		body, err := client.GetRequestBody("", nil)
		if err != nil {
			t.Errorf("Not expecting any error with nil body")
		}
		if body != nil {
			t.Errorf("Not expecting any body in return with nil body passed")
		}
	})

	//SubTest for Testing JSON Body
	t.Run("TestGetRequestBodyWithJSON", func(t *testing.T) {
		requestBody := []string{"foo", "bar"}
		actualBody, err := client.GetRequestBody("application/json", requestBody)
		if err != nil {
			t.Errorf("No error during marshalling of request body %v", requestBody)
		}
		expectedBody, _ := json.Marshal(requestBody)
		if string(expectedBody) != string(actualBody) {
			t.Errorf("Expected Body %v is not equal to the actual body: %v", string(expectedBody), string(actualBody))
		}
	})

	//SubTest for Testing Default Body without content-type
	t.Run("TestGetRequestBodyWithJSON", func(t *testing.T) {
		requestBody := []string{"foo", "bar"}
		actualBody, err := client.GetRequestBody("", requestBody)
		if err != nil {
			t.Errorf("No error during marshalling of request body %v", requestBody)
		}
		expectedBody, _ := json.Marshal(requestBody)
		if string(expectedBody) != string(actualBody) {
			t.Errorf("Expected Body %v is not equal to the actual body: %v", string(expectedBody), string(actualBody))
		}
	})
}
