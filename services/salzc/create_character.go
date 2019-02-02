package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/parnurzeal/gorequest"
	"github.com/schweigert/mga/libraries/randomizer"
)

func createCharacterRequest() (gorequest.Response, string, []error) {
	return gorequest.New().Post(SalzWebCharacterPath).Send(UserCharacter).End()
}

func createCharacter() {
	UserCharacter.AccountID = UserAccount.ID
	UserCharacter.Name = fmt.Sprintf("char#%s@%s", strconv.Itoa(randomizer.Int(100000)), UserAccount.Username)

	resp, body, errs := createCharacterRequest()
	if len(errs) > 0 {
		panic(errs)
	}

	json.Unmarshal([]byte(body), &UserCharacter)

	if resp.StatusCode != 200 {
		panic("Status error")
	}
}
