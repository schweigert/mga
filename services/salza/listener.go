package main

import (
	"errors"
	"os"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/schweigert/mga/libraries/db"
	"github.com/schweigert/mga/libraries/randomizer"
	"github.com/schweigert/mga/model"
)

// Listener server
type Listener struct{}

// Auth an accountg
func (listener *Listener) Auth(account model.Account, response *model.Account) (err error) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	dbc.Find(response, "id = ?", account.ID)

	if account.Password == response.Password && account.ID != 0 && response.ID == account.ID {
		*response = account
		response.AuthToken = strconv.Itoa(randomizer.Int(10000000))

		client := redis.NewClient(&redis.Options{
			Addr: os.Getenv("REDIS_ADDR"),
			DB:   0,
		})

		defer func() {
			errO := client.Close()
			if errO != nil {
				panic(errO)
			}
		}()

		err = client.Set(account.AuthKey(), response.AuthToken, 0).Err()
		if err != nil {
			panic(err)
		}

		return nil
	} else {
		*response = model.Account{}
		return errors.New("auth error")
	}
}
