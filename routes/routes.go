package routes

import (
	"task0/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/me", handlers.GetUserInfo)
}