package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/schweigert/mga/model"
	"github.com/schweigert/mga/services/rudydb/internal/db"
)

func IndexItemHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	items := &[]model.Item{}

	dbc.Find(items)
	c.JSON(200, items)
}

func ShowItemHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	item := &model.Item{}

	if c.BindJSON(item) == nil {
		dbc.Where("id = ?", item.ID).First(item)
		c.JSON(200, item)
	}
}

func CreateItemHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	item := &model.Item{}

	if c.BindJSON(item) == nil {
		c.JSON(200, dbc.Create(item))
	}
}

func UpdateItemHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	item := &model.Item{}

	if c.BindJSON(item) == nil {
		c.JSON(200, dbc.Save(item))
	}
}

func DestroyItemHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	item := &model.Item{}

	if c.BindJSON(item) == nil {
		c.JSON(200, dbc.Delete(item))
	}
}
