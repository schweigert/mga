package main

import (
	"encoding/json"

	"github.com/parnurzeal/gorequest"
)

func createAccountRequest() (gorequest.Response, string, []error) {
	return gorequest.New().Post(SalzWebAccountPath).Send(UserAccount).End()
}

func createAccount() {
	resp, body, errs := createAccountRequest()
	if len(errs) > 0 {
		panic(errs)
	}

	json.Unmarshal([]byte(body), &UserAccount)

	if resp.StatusCode != 200 {
		panic("Status error")
	}
}
