package model

// Character store positions and combat data of an character
type Character struct {
	Base

	PositionX  int
	PositionY  int
	DirectionX int
	DirectionY int
	Region     int

	Name  string
	HP    int
	MaxHP int
	Atack int

	AuthToken string

	Items     []Item
	AccountID uint
}

// Move to direction position
func (character *Character) Move() {
	character.PositionX += character.DirectionX
	character.PositionY += character.DirectionY

	character.Region %= 5
	character.PositionX %= 100
	character.PositionY %= 100

	if character.PositionX < 0 {
		character.PositionX *= -1
	}

	if character.PositionY < 0 {
		character.PositionY *= -1
	}

	character.DirectionX = 0
	character.DirectionY = 0
}
