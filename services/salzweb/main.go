package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/schweigert/mga/libraries/db"
	"github.com/schweigert/mga/model"
)

var router *gin.Engine

func migrate() {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	dbc.AutoMigrate(
		&model.Account{},
		&model.Character{},
		&model.Item{},
	)
}

func init() {
	migrate()

	router = gin.Default()
}

func crateAccountHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	account := &model.Account{}
	if c.BindJSON(account) == nil {
		dbc.Create(account)
		c.JSON(200, account)
	}
}

func crateCharacterHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	character := &model.Character{}
	if c.BindJSON(character) == nil {
		dbc.Create(character)
		c.JSON(200, character)
	}
}

func main() {
	router.POST("account/create", crateAccountHandler)
	router.POST("character/create", crateCharacterHandler)
	log.Panic(router.Run(os.Getenv("ADDR")))
}
