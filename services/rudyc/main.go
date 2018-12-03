package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/schweigert/mga/libraries/metric"
	"github.com/schweigert/mga/libraries/randomizer"
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
	USER_ACCOUNT.Username = fmt.Sprintf("user%s", strconv.Itoa(randomizer.Int(100000)))
}

func initPassword() {
	unsecuredPassword := fmt.Sprintf("pass%s", strconv.Itoa(randomizer.Int(100000)))
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
}

func main() {
	time.Sleep(10 * time.Second)
	log.Println("Starting Rudy Client...")
	log.Println("Username:", USER_ACCOUNT.Username)
	log.Println("Password:", USER_ACCOUNT.Password)

	log.Println("Rudyweb URL:", RUDYWEB_URL)

	steps()
}
