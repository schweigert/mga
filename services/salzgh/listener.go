package main

import (
	"errors"
	"net/rpc"
	"os"
	"sync"

	"github.com/schweigert/mga/model"
)

type Listener struct {
	CharacterList []model.Character
	Map           [5][100][100][]model.Character
	mapOperation  sync.Mutex
	Tokens        map[string]string
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

func NewListener() *Listener {
	return &Listener{Tokens: make(map[string]string)}
}

func (listener *Listener) checkAccount(account model.Account) bool {
	cachedToken, ok := listener.Tokens[account.Username]
	if ok {
		if account.CanAuth(cachedToken) {
			return true
		}
	}

	rpc, err := rpc.Dial("tcp", os.Getenv("SALZA_ADDR"))
	if err != nil {
		panic(err)
	}

	check := false

	err = rpc.Call("Listener.Check", account, &check)
	if err != nil {
		panic(err)
	}

	if check {
		listener.Tokens[account.Username] = account.AuthToken
	}

	return check
}

func (listener *Listener) SelectCharacter(account model.Account, selectedCharacter *model.Character) (err error) {
	if !listener.checkAccount(account) {
		return errors.New("not authenticated")
	}

	if len(account.Characters) <= 0 {
		return errors.New("Expected an character")
	}

	selectedCharacter = &account.Characters[0]
	selectedCharacter.AuthToken = account.AuthToken

	listener.appendCharacter(account.Characters[0])
	listener.respawCharacterIntoMap(account.Characters[0])

	return nil
}

func (listener *Listener) RoiCharacter(character model.Character, characterRoi *[100][100][]model.Character) error {
	roi := listener.regionOfInterest(character)
	*characterRoi = roi
	return nil
}

// MoveCharacter transfer one character to your direction
func (listener *Listener) MoveCharacter(character model.Character, postCharacter *model.Character) error {
	if character.DirectionX < -1 || character.DirectionX > 1 {
		if character.DirectionY < -1 || character.DirectionY > 1 {
			return errors.New("DirectionN != (-1, 0, 1)")
		}
	}

	err := listener.moveCharacter(character)
	postCharacter = &character
	postCharacter.Move()
	return err
}
