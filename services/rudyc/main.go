package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/schweigert/mga/libraries/metric"
	"github.com/schweigert/mga/libraries/security"
	"github.com/schweigert/mga/model"
)

var (
	USER_ACCOUNT           model.Account
	USER_CHARACTER         model.Character
	RUDYWEB_URL            string
	RUDYWEB_ACCOUNT_PATH   string
	RUDYWEB_CHARACTER_PATH string
)

func initUsername() {
	USER_ACCOUNT.Username = "user#" + strconv.Itoa(rand.Int()%100000)
}

func initPassword() {
	unsecuredPassword := "pass#" + strconv.Itoa(rand.Int()%100000)
	USER_ACCOUNT.Password = security.Hash(unsecuredPassword)
}

func initRudywebUrl() {
	RUDYWEB_URL = os.Getenv("RUDYWEB_URL")
	RUDYWEB_ACCOUNT_PATH = RUDYWEB_URL + "/account/create"
	RUDYWEB_CHARACTER_PATH = RUDYWEB_URL + "/character/create"
}

func init() {
	initUsername()
	initPassword()
	initRudywebUrl()
}

func steps() {
	metric.Timer("create_account", createAccount)
	metric.Timer("create_character", createCharacter)

	log.Println(USER_CHARACTER)
}

func main() {
	log.Println("Starting Rudy Client...")
	log.Println("Username:", USER_ACCOUNT.Username)
	log.Println("Password:", USER_ACCOUNT.Password)

	log.Println("Rudyweb URL:", RUDYWEB_URL)

	steps()
}
