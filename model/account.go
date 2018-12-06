package model

import "fmt"

type Account struct {
	Base

	Username string
	Password string

	AuthToken string

	Characters []Character
}

func (account *Account) AuthKey() string {
	return fmt.Sprintf("authToken:%s", account.Username)
}

func (account *Account) CanAuth(otherToken string) bool {
	return account.AuthToken == otherToken
}
