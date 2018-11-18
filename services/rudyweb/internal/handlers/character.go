package handlers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	"github.com/schweigert/mga/model"
)

func CreateCharacterHandler(c *gin.Context) {
	character := &model.Character{}

	if c.BindJSON(character) == nil {
		resp, body, _ := gorequest.New().Post(os.Getenv("RUDYDB_ADDR") + "/character/create").Send(character).End()
		c.JSON(resp.StatusCode, &gin.H{
			"body": body,
		})
	}
}
