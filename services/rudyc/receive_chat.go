package main

import (
	"log"

	"github.com/schweigert/mga/model"
)

func receiveChats() {
	chatSet := model.ChatSet{}

	err := RPCClient.Call("Listener.ReceiveChat", UserCharacter, &chatSet)
	if err != nil {
		panic(err)
	}

	for _, chat := range chatSet.Set {
		log.Println(chat.Line)
	}
}
