package main

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/go-redis/redis"
	"github.com/schweigert/mga/libraries/metric"
	"github.com/schweigert/mga/libraries/randomizer"
	"github.com/schweigert/mga/model"
)

var autorizeCounterMutex sync.Mutex
var authorizedAccounts int

func sendAuthorizedAccountsToGrafana() {
	for {
		autorizeCounterMutex.Lock()
		defer autorizeCounterMutex.Unlock()

		metric.Int("go.accounts", authorizedAccounts)
	}
}

// Autorize an account to connect into rudy game architecture
func (l *Listener) Autorize(account model.Account, authAccount *model.Account) (err error) {
	go func() {
		autorizeCounterMutex.Lock()
		defer autorizeCounterMutex.Unlock()

		authorizedAccounts++
	}()

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
