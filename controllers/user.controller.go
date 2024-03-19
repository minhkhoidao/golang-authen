package controllers

import (
	"golang-authen/driver"
	"golang-authen/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(ctx *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not hash password"})
		return
	}

	user := models.User{
		Email:    body.Email,
		Password: string(hash),
	}
	result := driver.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not create user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
