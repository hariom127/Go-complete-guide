package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/apis/models"
	"example.com/apis/utils"
	"github.com/gin-gonic/gin"
)

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
	token := context.Request.Header.Get("Authorization")
	fmt.Println("token===", token)

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Token is required"})
		return
	}
	userId, err := utils.VerifyToken(token)
	fmt.Println("tokenData===", userId)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	fmt.Println("token===", token)

	var event models.Event
	err = context.ShouldBindBodyWithJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request body"})
		return
	}
	event.UserId = userId

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request body", "err": err})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event has been created !", "event": event})

}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
	}

	_, err = models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
	}

	var updateEvent models.Event

	err = context.ShouldBindJSON(&updateEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
	}

	updateEvent.ID = id
	err = updateEvent.Update()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to update event."})
	}
	context.JSON(http.StatusAccepted, gin.H{"message": "Event has been updated successfully"})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Event not found"})
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete"})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event has been delete successfully"})

}
