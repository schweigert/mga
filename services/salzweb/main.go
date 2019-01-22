package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

func main() {
	log.Panic(router.Run(os.Getenv("ADDR")))
}
