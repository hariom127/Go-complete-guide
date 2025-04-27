package routes

import (
	"example.com/apis/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// these function getEvents, createEvents are available here because its part of same package event its not started with uppercase later
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authentication)
	authenticated.POST("/events", middleware.Authentication, createEvents)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
}
