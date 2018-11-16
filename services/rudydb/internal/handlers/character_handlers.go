package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/schweigert/mga/model"
	"github.com/schweigert/mga/services/rudydb/internal/db"
)

func IndexCharacterHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	characters := &[]model.Character{}

	dbc.Find(characters)
	c.JSON(200, characters)
}

func ShowCharacterHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	id := c.Param("id")
	character := &model.Character{}

	dbc.Where("id = ?", id).First(character)
	c.JSON(200, character)
}

func CreateCharacterHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	character := &model.Character{}

	if c.BindJSON(character) == nil {
		c.JSON(200, dbc.Create(character))
	}
}
