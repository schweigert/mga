package main

import (
	"errors"
	"strconv"

	"github.com/schweigert/mga/libraries/db"
	"github.com/schweigert/mga/libraries/randomizer"
	"github.com/schweigert/mga/model"
)

// Auth an accountg
func (listener *Listener) Auth(account model.Account, response *model.Account) (err error) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	dbc.Find(response, "id = ?", account.ID)

	if account.Password == response.Password && account.ID != 0 && response.ID == account.ID {
		*response = account
		response.AuthToken = strconv.Itoa(randomizer.Int(10000000))

		err = client.Set(account.AuthKey(), response.AuthToken, 0).Err()
		if err != nil {
			panic(err)
		}

		return nil
	}

	*response = model.Account{}
	return errors.New("auth error")
}
