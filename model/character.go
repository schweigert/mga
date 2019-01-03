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
