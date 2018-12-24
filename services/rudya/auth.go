package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/parnurzeal/gorequest"
	"github.com/schweigert/mga/model"
)

// Auth receive an account and return true or false if can auth
func (l *Listener) Auth(account model.Account, canAuth *bool) error {
	var status bool
	canAuth = &status

	resp, body, _ := gorequest.New().Post(os.Getenv("RUDYDB_ADDR") + "/account/show").Send(account).End()

	if resp.StatusCode != 200 {
		panic("Status code != 200")
	}

	dbAccount := &model.Account{}
	err := json.Unmarshal([]byte(body), dbAccount)
	if err != nil {
		panic(err)
	}

	if dbAccount.Password != account.Password {
		return errors.New("Auth error")
	}

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

	token, err := client.Get(account.AuthKey()).Result()

	if err != nil {
		status = false
	} else {
		status = account.CanAuth(token)
	}

	log.Println("Can auth for account", account.ID, ":", *canAuth)

	return nil
}
