package main

import "github.com/schweigert/mga/model"

// Listener server
type Listener struct {
	auths map[string]string
}

// Auth an account
func (listener *Listener) Auth(account model.Account, response *bool) (err error) {
	*response = true

	return
}
