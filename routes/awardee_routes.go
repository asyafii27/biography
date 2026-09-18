package routes

import (
	"biography-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAwardeeRoutes(routerGroup *gin.RouterGroup, controller *controllers.AwardeeController) {
	routes := routerGroup.Group("/awardees")
	{
		routes.GET("", controller.GetAll)
		routes.GET("/:id", controller.GetByID)
		routes.POST("", controller.Create)
		routes.PUT("/:id", controller.Update)
		routes.DELETE("/:id", controller.Delete)
	}
}
