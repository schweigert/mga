package main

import "log"

func selectCharacter() {
	err := RPCClient.Call("Listener.SelectCharacter", &UserAccount, &UserCharacter)
	if err != nil {
		panic(err)
	}

	log.Println("UserAccount:", UserAccount)
	log.Println("UserCharacter:", UserCharacter)
}
