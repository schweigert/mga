package main

import (
	"log"

	"github.com/schweigert/mga/model"
)

// Listener over TCP
type Listener struct {
	channels map[uint]chan model.Chat
}

// NewListener object
func NewListener() *Listener {
	return &Listener{channels: make(map[uint]chan model.Chat)}
}

// SendChat to others players
func (listener *Listener) SendChat(chat model.Chat, ret *model.Chat) (err error) {
	_, ok := listener.channels[chat.Character.ID]

	if !ok {
		listener.channels[chat.Character.ID] = make(chan model.Chat, 100)
	}

	for id, channel := range listener.channels {
		if chat.Character.ID != id {
			channel <- chat
		}
	}

	log.Println(chat.Character.ID, "|>", chat.Line)
	return nil
}

// ReceiveChat send a list of messages to the chat
func (listener *Listener) ReceiveChat(character model.Character, ret *model.ChatSet) error {
	messages := []model.Chat{}

	channel, ok := listener.channels[character.ID]

	if !ok {
		return nil
	}

	for {
		select {
		case chat := <-channel:
			messages = append(messages, chat)
		default:
			ret.Set = messages
			log.Println("Writing", len(ret.Set), "for", character.ID)
			return nil
		}
	}
}
