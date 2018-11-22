package main

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/schweigert/mga/libraries/security"
	"github.com/schweigert/mga/model"
)

var USER_ACCOUNT model.Account

func initUsername() {
	USER_ACCOUNT.Username = "user#" + strconv.Itoa(rand.Int()%100000)
}

func initPassword() {
	unsecuredPassword := "pass#" + strconv.Itoa(rand.Int()%100000)
	USER_ACCOUNT.Password = security.Hash(unsecuredPassword)
}

func init() {
	initUsername()
	initPassword()
}

func main() {
	log.Println("Starting Rudy Client...")
	log.Println("Username:", USER_ACCOUNT.Username)
	log.Println("Password:", USER_ACCOUNT.Password)
}
