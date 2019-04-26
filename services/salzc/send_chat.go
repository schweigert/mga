package main

import (
	"fmt"
	"log"

	"github.com/schweigert/mga/model"
)

func sendChat() {
	defer chatWaitGroup.Done()
	chat := model.Chat{}
	chat.Line = fmt.Sprintf("%s: I'm in this position: (%3d, %3d)", UserCharacter.Name, UserCharacter.PositionX, UserCharacter.PositionY)
	chat.Character = UserCharacter

	log.Println(chat, &chat)

	err := RPCChatClientSender.Call("Listener.SendChat", chat, &chat)
	if err != nil {
		panic(err)
	}

	log.Println("Chat submited")
}
