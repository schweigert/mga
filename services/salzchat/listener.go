package main

import (
	"log"

	"github.com/schweigert/mga/model"
)

type Listener struct {
}

func NewListener() *Listener {
	return &Listener{}
}

func (listener *Listener) SendChat(chat model.Chat, ret *model.Chat) (err error) {
	log.Println(chat.Character.ID, "|>", chat.Line)
	return nil
}
