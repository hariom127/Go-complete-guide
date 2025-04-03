package routes

import (
	"fmt"
	"net/http"

	"example.com/apis/models"
	"example.com/apis/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save user."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User has been created successfully."})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse user."})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	fmt.Println("user===", user)
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate the user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successfully", "token": token})

}
