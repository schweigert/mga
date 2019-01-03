package main

import (
	"errors"

	"github.com/schweigert/mga/model"
)

// MoveCharacter transfer one character to your direction
func (l *Listener) MoveCharacter(character model.Character, postCharacter *model.Character) error {
	if character.DirectionX < -1 || character.DirectionX > 1 {
		if character.DirectionY < -1 || character.DirectionY > 1 {
			return errors.New("DirectionN != (-1, 0, 1)")
		}
	}

	err := l.moveCharacter(character)
	postCharacter = &character
	postCharacter.Move()
	return err
}
