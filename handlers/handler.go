package handlers

import (
	"log"
	"net/http"
	"task0/clients"
	"task0/dto"
	"task0/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(ctx *gin.Context) {
	// Populate the relevant fields

	user := models.User{
		Email: "iJeddy5.0@gmail.com",
		Name:  "Iwegbu Jedidiah",
		Stack: "Go/Gin",
	}

	timestamp := time.Now().UTC().Format(time.RFC3339)

	factData, err := clients.GetKittyQuote()
	if err != nil {
		log.Println("Sad meow. Failed to fetch kitty quote of the day", err)
		factData = models.NekoQuote{Fact: "Could not fetch cat fact at the moment."}
	}

	response := dto.ResponseDto{
		Status:    "success",
		User:      user,
		Timestamp: timestamp,
		Fact:      factData.Fact,
	}
	ctx.JSON(http.StatusOK, response)

}
