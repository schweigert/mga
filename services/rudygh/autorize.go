package main

import (
	"net/rpc"
	"os"

	"github.com/schweigert/mga/model"
)

func (l *Listener) Autorize(account model.Account, authAccount *model.Account) error {
	client, err := rpc.Dial("tcp", os.Getenv("RUDYA_ADDR"))
	if err != nil {
		panic(err)
	}

	defer client.Close()
	err = client.Call("Listener.Autorize", account, &authAccount)
	if err != nil {
		panic(err)
	}

	return nil
}
