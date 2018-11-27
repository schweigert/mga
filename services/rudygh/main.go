package main

import (
	"log"
	"net"
	"net/rpc"
)

func main() {
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:4321")
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}

	listener := &Listener{}
	rpc.Register(listener)

	rpc.Accept(inbound)
}
