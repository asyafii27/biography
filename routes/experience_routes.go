package routes

import (
	"biography-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterExperienceRoutes(routerGroup *gin.RouterGroup, controller *controllers.ExperienceController) {
	expRoutes := routerGroup.Group("/experiences")
	{
		expRoutes.GET("", controller.GetAll)
		expRoutes.GET("/:id", controller.GetByID)
		expRoutes.POST("", controller.Create)
		expRoutes.PUT("/:id", controller.Update)
		expRoutes.DELETE("/:id", controller.Delete)
	}
}
