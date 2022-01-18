package main

import (
	"fmt"
	"net/http"

	"github.com/hemuku90/http-client-go/gohttp"
)

const (
	GithubURL = "https://api.github.com"
	GetURL    = "http://localhost:8080/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	PostUrl   = "http://localhost:8080/v1/organisation/accounts"
	DeleteURL = "http://localhost:8080/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc?version=0"
)

//Singleton
var (
	httpClient = gohttp.NewBuilder().
		Build()
)

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

//Main
func main() {
	fmt.Println(" ########### POST Request Output #######")
	postRequest()
	fmt.Println(" ########### GET Request Output #######")
	getRequest()
	fmt.Println(" ########### DELETE Request Output #######")
	deleteRequest()
}

func getRequest() {
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")

	response, err := httpClient.Get(PostUrl, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.String(), response.GetStatusCode())
}
func deleteRequest() {
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	response, err := httpClient.Delete(DeleteURL, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.String(), response.GetStatusCode())
}

func postRequest() {
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	body := payload()
	response, err := httpClient.Post(PostUrl, nil, body)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.String(), response.GetStatusCode())
}
