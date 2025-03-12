package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not save the user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created!"})

}

