package main

import (
	"net/rpc"
	"os"
)

func CheckAccount() {
	var err error
	defer RPCClient.Close()

	RPCClient, err = rpc.Dial("tcp", os.Getenv("SALZA_ADDR"))
	if err != nil {
		panic(err)
	}

	check := false

	err = RPCClient.Call("Listener.Check", UserAccount, &check)
	if err != nil {
		panic(err)
	}

	if !check {
		panic("User not checked")
	}
}
