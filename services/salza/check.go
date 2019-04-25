package main

import (
	"github.com/schweigert/mga/model"
)

// Check server
func (listener *Listener) Check(account *model.Account, check *bool) (err error) {
	token, err := client.Get(account.AuthKey()).Result()
	if err != nil {
		panic(err)
	}

	*check = token == account.AuthToken

	return
}
