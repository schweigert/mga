package main

import (
	"github.com/gin-gonic/gin"
	"github.com/schweigert/mga/services/rudydb/internal/handlers"
)

func routes(router *gin.Engine) {
	router.GET("/")

	router.GET("/accounts", handlers.IndexAccountHandler)
	router.GET("/account/show/:id", handlers.ShowAccountHandler)
	router.POST("/account/create", handlers.CreateAccountHandler)
}
