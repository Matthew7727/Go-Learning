package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err :=strconv.ParseInt(context.Param("id"),10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"could not parse event ID"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"ahh fuck can't find id"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"failed to register"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "registered"})

}

func cancelRegistration(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err :=strconv.ParseInt(context.Param("id"),10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"could not parse event ID"})
		return
	}

	var event models.Event 
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"failed to cancel"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "cancelled"})


}