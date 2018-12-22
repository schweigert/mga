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

	err = client.Call("Listener.Autorize", UserAccount, &UserAccount)
	if err != nil {
		panic(err)
	}

	log.Println("AUTHED CLIENT:", UserAccount)
}
