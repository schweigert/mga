package main

import (
	"log"
	"net/rpc"
	"os"

	"github.com/schweigert/mga/model"
)

// Autorize generate an token for a account
func (l *Listener) Autorize(account model.Account, authAccount *model.Account) error {
	client, err := rpc.Dial("tcp", os.Getenv("RUDYA_ADDR"))
	if err != nil {
		panic(err)
	}

	defer func() {
		errO := client.Close()
		if errO != nil {
			panic(errO)
		}
	}()

	err = client.Call("Listener.Autorize", account, &authAccount)
	if err != nil {
		panic(err)
	}

	AuthTokenCache[int(account.ID)] = authAccount.AuthToken
	log.Println("Creating auth cache:", int(account.ID), authAccount.AuthToken)

	authAccount.Characters = account.Characters
	return nil
}
