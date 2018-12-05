package model

import "fmt"

type Account struct {
	Base

	Username string
	Password string

	Characters []Character
}

func (account *Account) AuthKey() string {
	return fmt.Sprintf("authToken:%s", account.Username)
}
