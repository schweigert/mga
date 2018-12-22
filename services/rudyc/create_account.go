package main

import (
	"encoding/json"

	"github.com/parnurzeal/gorequest"
	"github.com/schweigert/mga/model"
)

type createAccountRequestEnvelope struct {
	Body string `json:"body"`
}

type createAccountRequestSubEnvelope struct {
	Account *model.Account `json:"Value"`
}

func createAccountRequest() (gorequest.Response, string, []error) {
	return gorequest.New().Post(RudyWebAccountPath).Send(UserAccount).End()
}

func createAccount() {
	resp, body, errs := createAccountRequest()
	if len(errs) != 0 {
		panic(errs[0])
	}

	if resp.StatusCode != 200 {
		panic("Status Code != 200 on create account")
	}

	data := &createAccountRequestEnvelope{}
	if err := json.Unmarshal([]byte(body), data); err != nil {
		panic(err)
	}

	value := &createAccountRequestSubEnvelope{Account: &UserAccount}
	if err := json.Unmarshal([]byte(data.Body), &value); err != nil {
		panic(err)
	}
}
