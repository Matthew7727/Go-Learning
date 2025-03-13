package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
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

func login(context *gin.Context) {

	var user models.User 
	err := context.ShouldBindBodyWithJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request"})
		return
	}

	err = user.ValidateLogin()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldnt auth user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "logged in", "token": token})
}