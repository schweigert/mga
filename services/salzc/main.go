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
	SalzWebURL           string
	SalzWebAccountPath   string
	SalzWebCharacterPath string
	RPCClient            *rpc.Client
	RPCGHClient          *rpc.Client
	ROI                  [100][100][]model.Character
)

func initUsername() {
	UserAccount.Username = fmt.Sprintf("user%s", strconv.Itoa(randomizer.Int(100000)))
}

func initPassword() {
	unsecuredPassword := fmt.Sprintf("pass%s", strconv.Itoa(randomizer.Int(100000)))
	UserAccount.Password = security.Hash(unsecuredPassword)
}

func initSalzWebURL() {
	SalzWebURL = os.Getenv("SALZWEB_URL")
	SalzWebAccountPath = SalzWebURL + "/account/create"
	SalzWebCharacterPath = SalzWebURL + "/character/create"
}

func init() {
	initUsername()
	initPassword()

	initSalzWebURL()
}

func steps() {
	metric.Timer("salzc.create_account", createAccount)
	metric.Timer("salzc.create_character", createCharacter)
	metric.Timer("salzc.auth_account", AuthAccount)
	metric.Timer("salzc.check_account", CheckAccount)
	metric.Timer("salzc.select_character", selectCharacter)
}

func main() {
	time.Sleep(10 * time.Second)
	log.Println("Username:", UserAccount.Username)
	log.Println("Password:", UserAccount.Password)

	steps()
}
