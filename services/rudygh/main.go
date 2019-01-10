package main

import (
	"log"
	"net"
	"net/rpc"
	"os"

	"github.com/schweigert/mga/model"
)

var (
	// AuthTokenCache store all auth tokens
	AuthTokenCache map[int]string
	// ChatChain store chat messages
	ChatChain map[uint]chan model.Chat
)

func init() {
	initAuthTokenCache()
	initChatChain()
}

func initChatChain() {
	ChatChain = make(map[uint]chan model.Chat)
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
