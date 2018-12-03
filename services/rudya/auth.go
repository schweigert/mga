package main

import (
	"os"

	"github.com/go-redis/redis"
	"github.com/schweigert/mga/model"
)

func (l *Listener) Auth(account model.Account, ok *bool) (err error) {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
		DB:   0,
	})
	return
}
