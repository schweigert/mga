package main

import (
	"net/rpc"
	"os"

	"github.com/schweigert/mga/model"
)

func (l *Listener) Auth(account model.Account, ack *bool) error {
	client, err := rpc.Dial("tcp", os.Getenv("RUDYA_ADDR"))
	if err != nil {
		panic(err)
	}

	defer client.Close()

	return client.Call("Listener.Auth", account, ack)
}
