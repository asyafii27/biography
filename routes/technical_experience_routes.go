package routes

import (
	"biography-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterTechnicalExperienceRoutes(routerGroup *gin.RouterGroup, controller *controllers.TechnicalExperienceController) {
	routes := routerGroup.Group("/technical-experiences")
	{
		routes.GET("", controller.GetAll)
		routes.GET("/:id", controller.GetByID)
		routes.POST("", controller.Create)
		routes.PUT("/:id", controller.Update)
		routes.DELETE("/:id", controller.Delete)
	}
}
