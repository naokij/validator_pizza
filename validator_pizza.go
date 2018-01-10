package validator_pizza

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type EmailResponse struct {
	Status            int
	Email             string
	Mx                bool
	Disposable        bool
	Alias             bool
	DidYouMean        *string `json:"did_you_mean"`
	RemainingRequests int     `json:"remaining_requests"`
	Error             string
}

type DomainResponse struct {
	Status            int
	Domain            string
	Mx                bool
	Disposable        bool
	DidYouMean        *string `json:"did_you_mean"`
	RemainingRequests int     `json:"remaining_requests"`
	Error             string
}

func ValidateEmail(email string) (resp EmailResponse, err error) {
	url := fmt.Sprintf("https://www.validator.pizza/email/%s", url.QueryEscape(email))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Add("Cache-Control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}
	return
}

func ValidateDomain(domain string) (resp DomainResponse, err error) {
	url := fmt.Sprintf("https://www.validator.pizza/domain/%s", url.QueryEscape(domain))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Add("Cache-Control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}
	return
}
