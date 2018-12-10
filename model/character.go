package model

type Character struct {
	Base

	Name  string
	HP    int
	MaxHP int
	Atack int

	AuthToken string

	Items     []Item
	AccountID uint
}
