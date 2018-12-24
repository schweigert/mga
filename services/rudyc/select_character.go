package main

import (
	"github.com/schweigert/mga/model"
)

func selectCharacter() {
	if len(UserAccount.Characters) == 0 {
		UserAccount.Characters = []model.Character{UserCharacter}
	}

	err := RPCClient.Call("Listener.SelectCharacter", &UserAccount, &UserCharacter)
	if err != nil {
		panic(err)
	}

	UserCharacter.AuthToken = UserAccount.AuthToken
}
