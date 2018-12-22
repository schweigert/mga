package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
	"time"

	"github.com/schweigert/mga/libraries/metric"
	"github.com/schweigert/mga/libraries/randomizer"
	"github.com/schweigert/mga/libraries/security"
	"github.com/schweigert/mga/model"
)

// Global Client Variables
var (
	UserAccount          model.Account
	UserCharacter        model.Character
	RudyWebURL           string
	RudyWebAccountPath   string
	RudyWebCharacterPath string
	RPCClient            *rpc.Client
)

func initUsername() {
	UserAccount.Username = fmt.Sprintf("user%s", strconv.Itoa(randomizer.Int(100000)))
}

func initPassword() {
	unsecuredPassword := fmt.Sprintf("pass%s", strconv.Itoa(randomizer.Int(100000)))
	UserAccount.Password = security.Hash(unsecuredPassword)
}

func initRudyWebURL() {
	RudyWebURL = os.Getenv("RUDYWEB_URL")
	RudyWebAccountPath = RudyWebURL + "/account/create"
	RudyWebCharacterPath = RudyWebURL + "/character/create"
}

func init() {
	initUsername()
	initPassword()
	initRudyWebURL()
}

func steps() {
	metric.Timer("create_account", createAccount)
	metric.Timer("create_character", createCharacter)
	metric.Timer("auth_account", authAccount)
	metric.Timer("select_character", selectCharacter)
}

func main() {
	time.Sleep(10 * time.Second)
	log.Println("Username:", UserAccount.Username)
	log.Println("Password:", UserAccount.Password)

	steps()
}
