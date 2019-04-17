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
	RPCChatClient        *rpc.Client
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

func initRPCChat() {
	var err error
	RPCChatClient, err = rpc.Dial("tcp", os.Getenv("SALZCHAT_ADDR"))
	if err != nil {
		panic(err)
	}
}

func initRPCAuthService() {
	var err error
	RPCClient, err = rpc.Dial("tcp", os.Getenv("SALZA_ADDR"))
	if err != nil {
		panic(err)
	}
}

func initRPCGameHandler() {
	var err error
	RPCGHClient, err = rpc.Dial("tcp", os.Getenv("SALZGH_ADDR"))
	if err != nil {
		panic(err)
	}
}

func init() {
	initUsername()
	initPassword()

	initSalzWebURL()
	initRPCChat()
	initRPCGameHandler()
	initRPCAuthService()
}

func steps() {
	metric.Timer("salzc.create_account", createAccount)
	metric.Timer("salzc.create_character", createCharacter)
	metric.Timer("salzc.auth_account", AuthAccount)
	metric.Timer("salzc.check_account", CheckAccount)
	metric.Timer("salzc.select_character", selectCharacter)

	for {
		metric.Timer("salzc.check_account", CheckAccount)
		metric.Timer("salzc.send_chat", sendChat)
		metric.Timer("salzc.receive_chat", receiveChat)
		metric.Timer("salzc.roi", roiCharacter)
		metric.Timer("salzc.move", moveCharacter)
	}
}

func main() {
	time.Sleep(10 * time.Second)
	log.Println("Username:", UserAccount.Username)
	log.Println("Password:", UserAccount.Password)

	steps()
}
