package handlers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	"github.com/schweigert/mga/model"
)

func CreateAccountHandler(c *gin.Context) {
	account := &model.Account{}

	if c.BindJSON(account) == nil {
		resp, body, _ := gorequest.New().Post(os.Getenv("RUDYDB_ADDR") + "/account/create").Send(account).End()
		c.JSON(resp.StatusCode, &gin.H{
			"body": body,
		})
	}
}
