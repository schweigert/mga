package main

import (
	"os"

	"github.com/go-redis/redis"
	"github.com/schweigert/mga/model"
)

// Listener server
func (listener *Listener) Check(account *model.Account, check *bool) (err error) {
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
		panic(err)
	}

	*check = token == account.AuthToken

	return
}
