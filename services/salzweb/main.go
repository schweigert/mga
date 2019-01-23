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

func main() {
	router.POST("accounts/", crateAccountHandler)
	log.Panic(router.Run(os.Getenv("ADDR")))
}
