package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/schweigert/mga/libraries/db"
	"github.com/schweigert/mga/model"
)

// IndexItemHandler handler
func IndexItemHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	items := &[]model.Item{}

	dbc.Find(items)
	c.JSON(200, items)
}

// ShowItemHandler handler
func ShowItemHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	item := &model.Item{}

	if c.BindJSON(item) == nil {
		dbc.Where("id = ?", item.ID).First(item)
		c.JSON(200, item)
	}
}

// CreateItemHandler handler
func CreateItemHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	item := &model.Item{}

	if c.BindJSON(item) == nil {
		c.JSON(200, dbc.Create(item))
	}
}

// UpdateItemHandler handler
func UpdateItemHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	item := &model.Item{}

	if c.BindJSON(item) == nil {
		c.JSON(200, dbc.Save(item))
	}
}

// DestroyItemHandler handler
func DestroyItemHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	item := &model.Item{}

	if c.BindJSON(item) == nil {
		c.JSON(200, dbc.Delete(item))
	}
}
