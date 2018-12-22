package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/parnurzeal/gorequest"
	"github.com/schweigert/mga/libraries/randomizer"
	"github.com/schweigert/mga/model"
)

type createCharacterRequestEnvelope struct {
	Body string `json:"body"`
}

type createCharacterRequestSubEnvelope struct {
	Character *model.Character `json:"Value"`
}

func createCharacterRequest() (gorequest.Response, string, []error) {
	return gorequest.New().Post(RudyWebCharacterPath).Send(UserCharacter).End()
}

func createCharacter() {
	UserCharacter.AccountID = UserAccount.ID
	UserCharacter.Name = fmt.Sprintf("char#%s@%s", strconv.Itoa(randomizer.Int(100000)), UserAccount.Username)

	resp, body, errs := createCharacterRequest()
	if len(errs) != 0 {
		panic(errs[0])
	}

	if resp.StatusCode != 200 {
		panic("Status Code != 200 on create account")
	}

	data := &createCharacterRequestEnvelope{}
	if err := json.Unmarshal([]byte(body), data); err != nil {
		panic(err)
	}

	value := &createCharacterRequestSubEnvelope{Character: &UserCharacter}
	if err := json.Unmarshal([]byte(data.Body), &value); err != nil {
		panic(err)
	}
}
