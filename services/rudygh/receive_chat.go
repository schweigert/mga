package main

import (
	"log"

	"github.com/schweigert/mga/model"
)

// ReceiveChat send a list of messages to the chat
func (l *Listener) ReceiveChat(character model.Character, ret *model.ChatSet) error {
	var messages []model.Chat
	for {
		select {
		case chat := <-ChatChain[character.ID]:
			messages = append(messages, chat)
		default:
			ret.Set = messages
			log.Println("Writing", len(ret.Set), "for", character.ID)
			return nil
		}
	}
}
