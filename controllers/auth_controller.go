package controllers

import (
	"biography-api/helpers"
	"biography-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService}
}

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (c *AuthController) Register(ctx *gin.Context) {
	var input RegisterInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error(http.StatusBadRequest, "Invalid input data", err.Error()))
		return
	}

	user, err := c.authService.Register(input.Name, input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error(http.StatusBadRequest, "Failed to register user", err.Error()))
		return
	}

	userData := gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	}

	ctx.JSON(http.StatusCreated, helpers.Success(http.StatusCreated, "User registered successfully", userData))
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (c *AuthController) Login(ctx *gin.Context) {
	var input LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error(http.StatusBadRequest, "Invalid input data", err.Error()))
		return
	}

	token, err := c.authService.Login(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, helpers.Error(http.StatusUnauthorized, "Login failed", err.Error()))
		return
	}

	tokenData := gin.H{
		"token": token,
	}

	ctx.JSON(http.StatusOK, helpers.Success(http.StatusOK, "Login successful", tokenData))
}
