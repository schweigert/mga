package main

import (
	"log"
	"net"
	"net/rpc"
	"os"
)

var (
	// AuthTokenCache store all auth tokens
	AuthTokenCache map[int]string
)

func init() {
	initAuthTokenCache()
}

func initAuthTokenCache() {
	AuthTokenCache = make(map[int]string)
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
