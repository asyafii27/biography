package routes

import (
	"biography-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterExperienceRoutes(router *gin.Engine, controller *controllers.ExperienceController) {
	api := router.Group("/api")
	{
		api.GET("/experiences", controller.GetAll)
		api.GET("/experiences/:id", controller.GetByID)
		api.POST("/experiences", controller.Create)
		api.PUT("/experiences/:id", controller.Update)
		api.DELETE("/experiences/:id", controller.Delete)
	}
}
