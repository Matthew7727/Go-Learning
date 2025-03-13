package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes (server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)

	authed := server.Group("/")
	authed.Use(middlewares.Authenticate)
	authed.POST("/events", createEvent)
	authed.PUT("/events/:id", updateEvent)
	authed.DELETE("events/:id", deleteEvent)
	authed.POST("/events/:id/register", registerForEvent)
	authed.DELETE("/events/:id/register", cancelRegistration)

	
	server.POST("/signup", signup)
	server.POST("/login", login)

}

