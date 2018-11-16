package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/schweigert/mga/model"
	"github.com/schweigert/mga/services/rudydb/internal/db"
)

func IndexAccountHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	accounts := &[]model.Account{}

	dbc.Find(accounts)
	c.JSON(200, accounts)
}

func ShowAccountHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	account := &model.Account{}

	if c.BindJSON(account) == nil {
		dbc.Where("id = ?", account.ID).First(account)
		c.JSON(200, account)
	}
}

func CreateAccountHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	account := &model.Account{}

	if c.BindJSON(account) == nil {
		c.JSON(200, dbc.Create(account))
	}
}

func UpdateAccountHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	account := &model.Account{}

	if c.BindJSON(account) == nil {
		c.JSON(200, dbc.Save(account))
	}
}

func DestroyAccountHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	account := &model.Account{}

	if c.BindJSON(account) == nil {
		c.JSON(200, dbc.Delete(account))
	}
}
