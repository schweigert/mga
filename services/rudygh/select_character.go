package main

import (
	"errors"

	"github.com/schweigert/mga/model"
)

func (l *Listener) SelectCharacter(account model.Account, selectedCharacter *model.Character) error {
	if !AuthToken(account) {
		return errors.New("AuthError")
	}

	if len(account.Characters) <= 0 {
		return errors.New("Expected an character")
	}

	selectedCharacter = &account.Characters[0]
	selectedCharacter.AuthToken = account.AuthToken

	return nil
}
