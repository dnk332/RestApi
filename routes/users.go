package routes

import (
	"fmt"
	"net/http"

	"example.com/restapi/models"
	"example.com/restapi/utils"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println("error: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	err = user.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})
}

func login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successfully.", "token": token})
}
