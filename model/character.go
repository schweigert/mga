package model

type Character struct {
	Base

	PositionX int
	PositionY int

	Name  string
	HP    int
	MaxHP int
	Atack int

	AuthToken string

	Items     []Item
	AccountID uint
}
