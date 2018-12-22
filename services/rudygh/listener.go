package main

import (
	"sync"

	"github.com/schweigert/mga/model"
)

// Listener implement all global functions with thread-safe operations cross all rpc calls
type Listener struct {
	CharacterList []model.Character
	mapOperation  sync.Mutex
}

func (l *Listener) appendCharacter(character model.Character) {
	l.mapOperation.Lock()
	defer l.mapOperation.Unlock()
	l.CharacterList = append(l.CharacterList, character)
}
