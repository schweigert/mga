package model

type Account struct {
	Base

	Username string
	Password string

	Characters []Character
}
