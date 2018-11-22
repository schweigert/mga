package main

import (
	"log"
	"math/rand"
	"strconv"
)

var USER_ID string
var PASS_ID string

func initUserID() {
	USER_ID = "user#" + strconv.Itoa(rand.Int()%100000)
}

func initPassID() {
	PASS_ID = "pass#" + strconv.Itoa(rand.Int()%100000)
}

func init() {
	initUserID()
	initPassID()
}

func main() {
	log.Println("Starting Rudy Client...")
	log.Println("USER_ID:", USER_ID)
	log.Println("PASS_ID:", PASS_ID)
}
