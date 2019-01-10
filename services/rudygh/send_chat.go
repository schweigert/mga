package main

import (
	"log"

	"github.com/schweigert/mga/model"
)

// SendChat send a message to the chat
func (l *Listener) SendChat(chat model.Chat, ret *model.Chat) error {
	ret = &chat
	log.Println(ret.Character.ID, "|>", ret.Line)
	go l.appendChat(chat)
	return nil
}

func (l *Listener) appendChat(chat model.Chat) {
	roi := l.regionOfInterest(chat.Character)
	for _, line := range roi {
		for _, cell := range line {
			for _, character := range cell {
				ChatChain[character.ID] <- chat
			}
		}
	}
}
