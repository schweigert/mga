package main

import (
	"log"
	"net"
	"net/rpc"
	"os"
)

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
