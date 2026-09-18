package routes

import (
	"biography-api/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes mengumpulkan semua konfigurasi routing dan prefix
func SetupRoutes(
	router *gin.Engine, 
	experienceController *controllers.ExperienceController, 
	authController *controllers.AuthController, 
	biographyController *controllers.BiographyController,
	awardeeController *controllers.AwardeeController,
	organizationController *controllers.OrganizationController,
	skillController *controllers.SkillController,
	technicalExperienceController *controllers.TechnicalExperienceController,
) {
	// Definisikan prefix utama /api di sini
	apiGroup := router.Group("/api")

	// Daftarkan route-route yang ada ke dalam grup /api
	RegisterExperienceRoutes(apiGroup, experienceController)
	RegisterAuthRoutes(apiGroup, authController)
	RegisterBiographyRoutes(apiGroup, biographyController)
	RegisterAwardeeRoutes(apiGroup, awardeeController)
	RegisterOrganizationRoutes(apiGroup, organizationController)
	RegisterSkillRoutes(apiGroup, skillController)
	RegisterTechnicalExperienceRoutes(apiGroup, technicalExperienceController)
}
