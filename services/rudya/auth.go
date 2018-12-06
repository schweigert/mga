package main

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/go-redis/redis"
	"github.com/parnurzeal/gorequest"
	"github.com/schweigert/mga/model"
)

func (l *Listener) Auth(account model.Account, canAuth *bool) error {
	resp, body, _ := gorequest.New().Post(os.Getenv("RUDYDB_ADDR") + "/account/show").Send(account).End()

	if resp.StatusCode != 200 {
		panic("Status code != 200")
	}

	dbAccount := &model.Account{}
	err := json.Unmarshal([]byte(body), dbAccount)

	if dbAccount.Password != account.Password {
		return errors.New("Auth error")
	}

	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
		DB:   0,
	})
	defer client.Close()

	token, err := client.Get(account.AuthKey()).Result()

	var status bool

	if err != nil {
		status = false
	} else {
		status = account.CanAuth(token)
	}

	canAuth = &status

	return nil
}
