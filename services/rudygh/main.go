package main

import (
	"log"
	"net"
	"net/rpc"
	"os"

	"github.com/schweigert/mga/libraries/surroundings"
)

var (
	SURROUNDINGS   [5]*surroundings.Surroundings
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
	SURROUNDINGS[0] = surroundings.NewSurroundings()
	SURROUNDINGS[1] = surroundings.NewSurroundings()
	SURROUNDINGS[2] = surroundings.NewSurroundings()
	SURROUNDINGS[3] = surroundings.NewSurroundings()
	SURROUNDINGS[4] = surroundings.NewSurroundings()
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
