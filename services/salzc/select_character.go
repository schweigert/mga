package main

import (
	"log"
)

func selectCharacter() {
	err := RPCGHClient.Call("Listener.SelectCharacter", UserAccount, &UserCharacter)
	if err != nil {
		panic(err)
	}

	log.Println("Selected character", UserCharacter)
}
