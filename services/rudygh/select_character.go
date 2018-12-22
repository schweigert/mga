package main

import (
	"errors"
	"log"

	"github.com/schweigert/mga/model"
)

// SelectCharacter select and deploy an character
func (l *Listener) SelectCharacter(account model.Account, selectedCharacter *model.Character) error {
	log.Println(account)
	log.Println(AuthTokenCache)
	if !AuthToken(account) {
		return errors.New("AuthError")
	}

	if len(account.Characters) <= 0 {
		return errors.New("Expected an character")
	}

	selectedCharacter = &account.Characters[0]
	selectedCharacter.AuthToken = account.AuthToken

	l.appendCharacter(account.Characters[0])

	return nil
}
