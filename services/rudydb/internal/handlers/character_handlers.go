package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/schweigert/mga/libraries/db"
	"github.com/schweigert/mga/model"
)

// IndexCharacterHandler handler
func IndexCharacterHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	characters := &[]model.Character{}

	dbc.Find(characters)
	c.JSON(200, characters)
}

// ShowCharacterHandler handler
func ShowCharacterHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	character := &model.Character{}

	if c.BindJSON(character) == nil {
		dbc.Where("id = ?", character.ID).First(character)
		c.JSON(200, character)
	}
}

// CreateCharacterHandler handler
func CreateCharacterHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	character := &model.Character{}

	if c.BindJSON(character) == nil {
		c.JSON(200, dbc.Create(character))
	}
}

// UpdateCharacterHandler handler
func UpdateCharacterHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	character := &model.Character{}

	if c.BindJSON(character) == nil {
		c.JSON(200, dbc.Save(character))
	}
}

// DestroyCharacterHandler handler
func DestroyCharacterHandler(c *gin.Context) {
	dbc := db.Connect()
	defer db.SafeClose(dbc)

	character := &model.Character{}

	if c.BindJSON(character) == nil {
		c.JSON(200, dbc.Delete(character))
	}
}
