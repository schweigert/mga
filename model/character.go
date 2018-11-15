package model

type Character struct {
	Base

	Name  string
	HP    int
	MaxHP int
	Atack int

	Items     []Item
	AccountID uint
}
