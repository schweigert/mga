package main

import (
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/schweigert/mga/libraries/randomizer"
	"github.com/schweigert/mga/model"
)

func (l *Listener) Autorize(account model.Account, authAccount *model.Account) (err error) {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
		DB:   0,
	})

	defer client.Close()

	randToken := strconv.Itoa(randomizer.Int(10000000))

	err = client.Set(account.AuthKey(), randToken, 0).Err()
	if err != nil {
		panic(err)
	}

	authAccount.AuthToken = randToken
	authAccount.Characters = account.Characters
	log.Println(authAccount)

	return
}
