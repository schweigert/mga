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
	return gorequest.New().Post(RUDYWEB_CHARACTER_PATH).Send(USER_CHARACTER).End()
}

func createCharacter() {
	USER_CHARACTER.AccountID = USER_ACCOUNT.ID
	USER_CHARACTER.Name = fmt.Sprintf("char#%s@%s", strconv.Itoa(randomizer.Int(100000)), USER_ACCOUNT.Username)

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

	value := &createCharacterRequestSubEnvelope{Character: &USER_CHARACTER}
	if err := json.Unmarshal([]byte(data.Body), &value); err != nil {
		panic(err)
	}
}
