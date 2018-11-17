package main

import (
	"github.com/gin-gonic/gin"
	"github.com/schweigert/mga/services/rudydb/internal/handlers"
)

func routes(router *gin.Engine) {
	router.GET("/")

	router.GET("/accounts", handlers.IndexAccountHandler)
	router.POST("/account/show", handlers.ShowAccountHandler)
	router.POST("/account/create", handlers.CreateAccountHandler)
	router.POST("/account/update", handlers.UpdateAccountHandler)
	router.POST("/account/destroy", handlers.DestroyAccountHandler)

	router.GET("/characters", handlers.IndexCharacterHandler)
	router.POST("/character/show", handlers.ShowCharacterHandler)
	router.POST("/character/create", handlers.CreateCharacterHandler)
	router.POST("/character/update", handlers.UpdateCharacterHandler)
	router.POST("/character/destroy", handlers.DestroyCharacterHandler)
}
