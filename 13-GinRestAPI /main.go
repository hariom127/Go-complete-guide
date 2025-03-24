package main

import (
	"net/http"

	"example.com/apis/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvents)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	var events = models.GetAllEvents()
	context.JSON(http.StatusOK, gin.H{"message": "Hello !🥳", "events": events})
}

func createEvents(context *gin.Context) {

	var event models.Event

	err := context.ShouldBindBodyWithJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request body"})
		return
	}
	event.ID = 1
	event.UserId = 1

	context.JSON(http.StatusCreated, gin.H{"message": "Event has been created !", "event": event})

}
