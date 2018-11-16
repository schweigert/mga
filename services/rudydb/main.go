package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	log.Println("Starting rudydb...")
	// Database configuration
	log.Println("db:migrate")
	migrate()

	// Webserver configuration
	log.Println("routes:add")
	router = gin.Default()
	routes(router)
}

func main() {
	panic(router.Run(os.Getenv("ADDR")))
}
