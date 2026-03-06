package gohttp

import (
	"testing"
)

/*
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
*/

const (
	getURL    = "http://accountapi:8080/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	postURL   = "http://accountapi:8080/v1/organisation/accounts"
	deleteURL = "http://accountapi:8080/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc?version=0"
)

// Mock Payload structure
type Data struct {
	Data *AccountData `json:"data,omitempty"`
}

type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
}

type AccountAttributes struct {
	BankID              string   `json:"bank_id,omitempty"`
	BankIDCode          string   `json:"bank_id_code,omitempty"`
	BaseCurrency        string   `json:"base_currency,omitempty"`
	Bic                 string   `json:"bic,omitempty"`
	Country             *string  `json:"country,omitempty"`
	Name                []string `json:"name,omitempty"`
	ValidationType      string   `json:"validation_type,omitempty"`
	ReferenceMask       string   `json:"reference_mask,omitempty"`
	AcceptanceQualifier string   `json:"acceptance_qualifier,omitempty"`
}

//Create request Payload
func payload() *Data {
	country := "GB"
	body := &AccountData{
		ID:             "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts",
		Attributes: &AccountAttributes{
			BankID:       "400300",
			BankIDCode:   "GBDSC",
			BaseCurrency: "GBP",
			Bic:          "NWBKGB22",
			Country:      &country,
			Name:         []string{"Samantha Holder"},
			//	UserDefinedData: ,
		},
	}
	data := &Data{
		Data: body,
	}
	return data
}

//Check if int is contained in a slice of int
func contains(s []int, val int) bool {
	for _, v := range s {
		if v == val {
			return true
		}
	}

	return false
}
func TestPost(t *testing.T) {
	client := &httpClient{}
	client.builder = &clientBuilder{}
	var body *Data = payload() //Mock request payload
	t.Run("TestPostRequest", func(t *testing.T) {
		_, err := client.Post(postURL, nil, body)
		if err != nil {
			t.Errorf("Error while getting the response from %s", postURL)
		}
		// fmt.Println(response.String())
	})
	t.Run("TestPostStatusCode", func(t *testing.T) {
		response, err := client.Post(postURL, nil, body)
		if err != nil {
			t.Errorf("Error while getting the response from %s", postURL)
		}
		actualStatusCode := response.GetStatusCode()
		expectedStatusCode := []int{201, 409}
		if !(contains(expectedStatusCode, actualStatusCode)) {
			t.Errorf("Expeted status codes were %v but got %v", expectedStatusCode, actualStatusCode)
		}
	})
}

func TestGet(t *testing.T) {
	client := &httpClient{}
	client.builder = &clientBuilder{}
	t.Run("TestGETRequest", func(t *testing.T) {
		_, err := client.Get(getURL, nil)
		if err != nil {
			t.Errorf("Error while getting the response from %s", getURL)
		}
	})
	// fmt.Println(response.String())
	t.Run("TestGETStatusCode", func(t *testing.T) {
		response, err := client.Get(getURL, nil)
		if err != nil {
			t.Errorf("Error while getting the response from %s", postURL)
		}
		actualStatusCode := response.GetStatusCode()
		expectedStatusCode := []int{200}
		if !(contains(expectedStatusCode, actualStatusCode)) {
			t.Errorf("Expeted status codes were %v but got %v", expectedStatusCode, actualStatusCode)
		}
	})
}

func TestDelete(t *testing.T) {
	client := &httpClient{}
	client.builder = &clientBuilder{}
	t.Run("TestDeleteRequest", func(t *testing.T) {
		_, err := client.Delete(deleteURL, nil)
		if err != nil {
			t.Errorf("Error while getting the response from %s", getURL)
		}
	})
	// fmt.Println(response.String())
	t.Run("TestDeleteStatusCode", func(t *testing.T) {
		response, err := client.Delete(deleteURL, nil)
		if err != nil {
			t.Errorf("Error while getting the response from %s", postURL)
		}
		actualStatusCode := response.GetStatusCode()
		expectedStatusCode := []int{200}
		if !(contains(expectedStatusCode, actualStatusCode)) {
			t.Errorf("Expeted status codes were %v but got %v", expectedStatusCode, actualStatusCode)
		}
	})
}
