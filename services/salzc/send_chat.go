package main

import (
	"fmt"
	"log"

	"github.com/schweigert/mga/model"
)

func sendChat() {
	chat := model.Chat{}
	chat.Line = fmt.Sprintf("%s: I'm in this position: (%3d, %3d)", UserCharacter.Name, UserCharacter.PositionX, UserCharacter.PositionY)
	chat.Character = UserCharacter

	log.Println(chat, &chat)

	err := RPCChatClient.Call("Listener.SendChat", chat, &chat)
	if err != nil {
		panic(err)
	}
}
