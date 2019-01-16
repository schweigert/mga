package main

import (
	"errors"

	"github.com/schweigert/mga/model"
)

// RoiCharacter return the region of interest of an character
func (l *Listener) RoiCharacter(character model.Character, characterRoi *[100][100][]model.Character) error {
	if AuthCharacter(character) {
		roi := l.regionOfInterest(character)
		*characterRoi = roi
		return nil
	}
	return errors.New("AuthError")
}
