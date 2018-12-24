package main

import (
	"sync"

	"github.com/schweigert/mga/model"
)

// Listener implement all global functions with thread-safe operations cross all rpc calls
type Listener struct {
	CharacterList []model.Character
	Map           [5][100][100][]model.Character
	mapOperation  sync.Mutex
}

func (l *Listener) appendCharacter(character model.Character) {
	l.mapOperation.Lock()
	defer l.mapOperation.Unlock()
	l.CharacterList = append(l.CharacterList, character)
}

func (l *Listener) respawCharacterIntoMap(character model.Character) {
	l.mapOperation.Lock()
	defer l.mapOperation.Unlock()

	l.Map[0][character.PositionX%100][character.PositionX%100] = append(l.Map[0][character.PositionX%100][character.PositionX%100], character)
}
