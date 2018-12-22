package main

import (
	"log"
	"net"
	"net/rpc"
	"os"

	"github.com/schweigert/mga/libraries/surroundings"
)

var (
	// Surroundings are a map structure of this service
	Surroundings [5]*surroundings.Surroundings
	// AuthTokenCache store all auth tokens
	AuthTokenCache map[int]string
)

func init() {
	initSurroundings()
	initAuthTokenCache()
}

func initAuthTokenCache() {
	AuthTokenCache = make(map[int]string)
}

func initSurroundings() {
	Surroundings[0] = surroundings.NewSurroundings()
	Surroundings[1] = surroundings.NewSurroundings()
	Surroundings[2] = surroundings.NewSurroundings()
	Surroundings[3] = surroundings.NewSurroundings()
	Surroundings[4] = surroundings.NewSurroundings()
}

func main() {
	addy, err := net.ResolveTCPAddr("tcp", os.Getenv("ADDR"))
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}

	listener := &Listener{}
	if err = rpc.Register(listener); err != nil {
		panic(err)
	}
	rpc.Accept(inbound)
}
