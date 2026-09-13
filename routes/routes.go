package routes

import (
	"biography-api/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes mengumpulkan semua konfigurasi routing dan prefix
func SetupRoutes(router *gin.Engine, experienceController *controllers.ExperienceController, authController *controllers.AuthController) {
	// Definisikan prefix utama /api di sini
	apiGroup := router.Group("/api")

	// Daftarkan route-route yang ada ke dalam grup /api
	RegisterExperienceRoutes(apiGroup, experienceController)
	RegisterAuthRoutes(apiGroup, authController)
}
