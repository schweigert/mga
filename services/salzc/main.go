package main

import (
	"fmt"
	"log"
	"net/rpc"
	"strconv"
	"time"

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
	ROI                  [100][100][]model.Character
)

func initUsername() {
	UserAccount.Username = fmt.Sprintf("user%s", strconv.Itoa(randomizer.Int(100000)))
}

func initPassword() {
	unsecuredPassword := fmt.Sprintf("pass%s", strconv.Itoa(randomizer.Int(100000)))
	UserAccount.Password = security.Hash(unsecuredPassword)
}

func init() {
	initUsername()
	initPassword()
}

func steps() {
}

func main() {
	time.Sleep(10 * time.Second)
	log.Println("Username:", UserAccount.Username)
	log.Println("Password:", UserAccount.Password)

	steps()
}
