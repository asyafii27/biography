package routes

import (
	"biography-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(routerGroup *gin.RouterGroup, authController *controllers.AuthController) {
	authRoutes := routerGroup.Group("/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
	}
}
