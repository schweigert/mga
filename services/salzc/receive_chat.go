package main

import (
	"log"

	"github.com/schweigert/mga/model"
)

func receiveChat() {
	chatSet := model.ChatSet{}

	err := RPCChatClient.Call("Listener.ReceiveChat", UserCharacter, &chatSet)
	if err != nil {
		panic(err)
	}

	for _, chat := range chatSet.Set {
		log.Println(chat.Line)
	}
}
