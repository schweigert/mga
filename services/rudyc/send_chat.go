package main

import (
	"fmt"

	"github.com/schweigert/mga/model"
)

func sendPositionMessageChat() {
	chat := model.Chat{}
	chat.Line = fmt.Sprintf("%s: I'm in this position: (%3d, %3d)", UserCharacter.Name, UserCharacter.PositionX, UserCharacter.PositionY)
	chat.Character = UserCharacter

	err := RPCClient.Call("Listener.SendChat", chat, &chat)
	if err != nil {
		panic(err)
	}
}
