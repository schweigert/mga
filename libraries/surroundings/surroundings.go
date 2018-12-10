package surroundings

import "github.com/schweigert/mga/model"

type Surroundings struct {
	CharacterLayer [100][100][]model.Character
}

func NewSurroundings() *Surroundings {
	return &Surroundings{}
}
