package main

import (
	"errors"
	"net/rpc"
	"os"

	"github.com/schweigert/mga/model"
)

type Listener struct {
	Tokens map[string]string
}

func (listener *Listener) checkAccount(account model.Account) bool {
	cachedToken, ok := listener.Tokens[account.Username]
	if ok {
		if account.CanAuth(cachedToken) {
			return true
		}
	}

	rpc, err := rpc.Dial("tcp", os.Getenv("SALZA_ADDR"))
	if err != nil {
		panic(err)
	}

	check := false

	err = rpc.Call("Listener.Check", account, &check)
	if err != nil {
		panic(err)
	}

	if check {
		listener.Tokens[account.Username] = account.AuthToken
	}

	return check
}

func (listener *Listener) SelectCharcter(account model.Account, selectedCharacter *model.Character) (err error) {
	if !listener.checkAccount(account) {
		return errors.New("not authenticated")
	}

	return nil
}
