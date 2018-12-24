package main

import (
	"net/rpc"
	"os"

	"github.com/schweigert/mga/model"
)

// AuthToken certify if an account can auth into rudygh
func AuthToken(account model.Account) bool {
	if AuthTokenCache[int(account.ID)] == account.AuthToken {
		return true
	}

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

	var ret bool
	err = client.Call("Listener.Auth", account, &ret)
	if err != nil {
		panic(err)
	}

	return ret
}
