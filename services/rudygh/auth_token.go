package main

import (
	"net/rpc"
	"os"

	"github.com/schweigert/mga/model"
)

func AuthToken(account model.Account) bool {
	client, err := rpc.Dial("tcp", os.Getenv("RUDYA_ADDR"))
	if err != nil {
		panic(err)
	}

	defer client.Close()
	var ret bool
	err = client.Call("Listener.Auth", account, &ret)
	if err != nil {
		panic(err)
	}

	return ret
}
