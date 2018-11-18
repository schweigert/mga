package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	log.Println("Starting rudyweb...")
	router = gin.Default()
}

func main() {
	panic(router.Run(os.Getenv("ADDR")))
}
