package main

import (
	"github.com/gin-gonic/gin"
	"github.com/schweigert/mga/services/rudyweb/internal/handlers"
)

func routes(router *gin.Engine) {
	router.POST("/account/create", handlers.CreateAccountHandler)
	router.POST("/character/create", handlers.CreateCharacterHandler)
}
