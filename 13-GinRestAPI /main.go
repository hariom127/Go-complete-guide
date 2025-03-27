package main

import (
	"net/http"
	"strconv"

	"example.com/apis/db"
	"example.com/apis/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvents)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	var events, err = models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server error", "events": events})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Hello !🥳", "events": events})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
	}
	result, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event", "err": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event has been retrieve successfully", "event": result})

}

func createEvents(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindBodyWithJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request body"})
		return
	}
	// event.ID = 1
	event.UserId = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request body", "err": err})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event has been created !", "event": event})

}
