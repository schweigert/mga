package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	log.Println("Starting rudydb...")
	router = gin.Default()

	Routes(router)
}

func main() {
	router.Run(os.Getenv("ADDR"))
}
