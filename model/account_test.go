package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountAuthKey(t *testing.T) {
	account := Account{Username: "TestingUser"}
	assert.Equal(t, "authToken:TestingUser", account.AuthKey())
}

func TestAccountCanAuth(t *testing.T) {
	account := Account{AuthToken: "MyToken"}
	assert.True(t, account.CanAuth("MyToken"))

}
