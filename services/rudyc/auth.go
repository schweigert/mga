package main

import (
	"log"
	"net/rpc"
	"os"
)

func authAccount() {
	client, err := rpc.Dial("tcp", os.Getenv("RUDYGH_ADDR"))
	if err != nil {
		panic(err)
	}

	defer client.Close()

	var authOk bool

	err = client.Call("Listener.Auth", USER_ACCOUNT, &authOk)
	if err != nil {
		panic(err)
	}

	log.Println("authOk:", authOk)
}
