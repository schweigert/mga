package main

import (
	"log"
	"net"
	"net/rpc"
	"os"

	"github.com/go-redis/redis"
)

var client *redis.Client

func main() {
	client = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
		DB:   0,
	})

	addy, err := net.ResolveTCPAddr("tcp", os.Getenv("ADDR"))
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}

	if err = rpc.Register(&Listener{}); err != nil {
		panic(err)
	}
	rpc.Accept(inbound)
}
