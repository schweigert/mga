package main

import (
	"log"
	"net/rpc"
	"os"
)

func AuthAccount() {
	var err error
	RPCClient, err = rpc.Dial("tcp", os.Getenv("SALZA_ADDR"))
	if err != nil {
		panic(err)
	}

	err = RPCClient.Call("Listener.Autorize", UserAccount, &UserAccount)
	if err != nil {
		panic(err)
	}

	log.Println("AUTHED CLIENT:", UserAccount)
}
