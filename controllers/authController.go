package controllers

import (
	"net/http"
	"rest-api-note-taking/models"
	"rest-api-note-taking/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterInput struct {
	Name     string `json:"name" binding:"required`
	Username string `json:"username" binding:"required`
	Email    string `json:"email" binding:"required`
	Password string `json:"password" binding:"required`
}

type LoginInput struct {
	Username string `json:"username" binding:"required`
	Password string `json:"password" binding:"required`
}

func Login(ctx *gin.Context) {
	input := LoginInput{}
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := ctx.MustGet("db").(*gorm.DB)
	jwtToken, err := models.LoginValid(input.Username, input.Password, db)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Login success", "token": jwtToken})
}

func Register(ctx *gin.Context) {
	input := RegisterInput{}
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !utils.IsValidEmail(input.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email (make sure email format correct)"})
		return
	}
	newUser := models.User{
		Name:     input.Name,
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}
	db := ctx.MustGet("db").(*gorm.DB)
	err := newUser.SaveUser(db)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := map[string]string{
		"name":     input.Name,
		"username": input.Username,
		"email":    input.Email,
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Registration success", "user": result})
}
