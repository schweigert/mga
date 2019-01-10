package main

import (
	"errors"

	"github.com/schweigert/mga/model"
)

// SelectCharacter select and deploy an character
func (l *Listener) SelectCharacter(account model.Account, selectedCharacter *model.Character) error {
	if !AuthToken(account) {
		return errors.New("AuthError")
	}

	if len(account.Characters) <= 0 {
		return errors.New("Expected an character")
	}

	selectedCharacter = &account.Characters[0]
	selectedCharacter.AuthToken = account.AuthToken

	l.appendCharacter(account.Characters[0])
	l.respawCharacterIntoMap(account.Characters[0])

	ChatChain[account.Characters[0].ID] = make(chan model.Chat, 20)

	return nil
}
