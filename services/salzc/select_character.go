package main

import (
	"log"
	"net/rpc"
	"os"
)

func selectCharacter() {
	var err error
	RPCGHClient, err = rpc.Dial("tcp", os.Getenv("SALZGH_ADDR"))
	if err != nil {
		panic(err)
	}

	err = RPCGHClient.Call("Listener.SelectCharacter", UserAccount, &UserCharacter)
	if err != nil {
		panic(err)
	}

	log.Println("Selected character", UserCharacter)
}
