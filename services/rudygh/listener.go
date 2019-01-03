package main

import (
	"errors"
	"sync"

	"github.com/schweigert/mga/model"
)

// Listener implement all global functions with thread-safe operations cross all rpc calls
type Listener struct {
	CharacterList []model.Character
	Map           [5][100][100][]model.Character
	mapOperation  sync.Mutex
}

func (l *Listener) moveCharacter(character model.Character) error {
	l.mapOperation.Lock()
	defer l.mapOperation.Unlock()
	// Remove old character
	for index, otherCharacter := range l.Map[character.Region][character.PositionX][character.PositionY] {
		if otherCharacter.ID == character.ID {
			characterArray := l.Map[character.Region][character.PositionX][character.PositionY]

			l.Map[character.Region][character.PositionX][character.PositionY] = append(characterArray[:index], characterArray[index+1:]...)

			character.Move()

			// Append to the new position
			l.Map[character.Region][character.PositionX][character.PositionY] = append(l.Map[character.Region][character.PositionX][character.PositionY], character)

			for index, otherListCharacter := range l.CharacterList {
				if otherListCharacter.ID == character.ID {
					l.CharacterList[index] = character
					return nil
				}
			}
			return errors.New("moveCharacter: Character not found into CharacterList")
		}
	}

	return errors.New("moveCharacter: Can find this character in this map")
}

func (l *Listener) regionOfInterest(character model.Character) [100][100][]model.Character {
	l.mapOperation.Lock()
	defer l.mapOperation.Unlock()

	return l.Map[character.Region%5]
}

func (l *Listener) appendCharacter(character model.Character) {
	l.mapOperation.Lock()
	defer l.mapOperation.Unlock()
	l.CharacterList = append(l.CharacterList, character)
}

func (l *Listener) respawCharacterIntoMap(character model.Character) {
	l.mapOperation.Lock()
	defer l.mapOperation.Unlock()

	character.Region = character.Region % 5
	character.PositionX = character.PositionX % 100
	character.PositionY = character.PositionY % 100

	l.Map[character.Region][character.PositionX][character.PositionY] = append(l.Map[character.Region][character.PositionX][character.PositionY], character)
	l.CharacterList = append(l.CharacterList, character)
}
