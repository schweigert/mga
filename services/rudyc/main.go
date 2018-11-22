package main

import (
	"crypto/sha1"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/schweigert/mga/model"
)

var USER_ACCOUNT model.Account

func initUsername() {
	USER_ACCOUNT.Username = "user#" + strconv.Itoa(rand.Int()%100000)
}

func initPassword() {
	USER_ACCOUNT.Password = "pass#" + strconv.Itoa(rand.Int()%100000)

	// Apply sha1 over this account
	USER_ACCOUNT.Password = fmt.Sprintf("%x", sha1.Sum([]byte(USER_ACCOUNT.Password)))
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
