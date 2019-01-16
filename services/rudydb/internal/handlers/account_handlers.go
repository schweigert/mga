package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/schweigert/mga/libraries/db"
	"github.com/schweigert/mga/model"
)

// IndexAccountHandler handler
func IndexAccountHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	accounts := &[]model.Account{}

	dbc.Preload("Characters").Find(accounts)
	c.JSON(200, accounts)
}

// ShowAccountHandler handler
func ShowAccountHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	account := &model.Account{}

	if c.BindJSON(account) == nil {
		dbc.Where("id = ?", account.ID).First(account)
		c.JSON(200, account)
	}
}

// CreateAccountHandler handler
func CreateAccountHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	account := &model.Account{}

	if c.BindJSON(account) == nil {
		c.JSON(200, dbc.Create(account))
	}
}

// UpdateAccountHandler handler
func UpdateAccountHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	account := &model.Account{}

	if c.BindJSON(account) == nil {
		c.JSON(200, dbc.Save(account))
	}
}

// DestroyAccountHandler handler
func DestroyAccountHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	account := &model.Account{}

	if c.BindJSON(account) == nil {
		c.JSON(200, dbc.Delete(account))
	}
}
