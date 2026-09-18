package routes

import (
	"biography-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterOrganizationRoutes(routerGroup *gin.RouterGroup, controller *controllers.OrganizationController) {
	routes := routerGroup.Group("/organizations")
	{
		routes.GET("", controller.GetAll)
		routes.GET("/:id", controller.GetByID)
		routes.POST("", controller.Create)
		routes.PUT("/:id", controller.Update)
		routes.DELETE("/:id", controller.Delete)
	}
}
