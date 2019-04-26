package main

import (
	"log"

	"github.com/schweigert/mga/model"
)

func receiveChat() {
	defer chatWaitGroup.Done()
	chatSet := model.ChatSet{}

	err := RPCChatClientReceiver.Call("Listener.ReceiveChat", UserCharacter, &chatSet)
	if err != nil {
		panic(err)
	}

	for _, chat := range chatSet.Set {
		log.Println(chat.Line)
	}
}
