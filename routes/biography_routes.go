package routes

import (
	"biography-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterBiographyRoutes(routerGroup *gin.RouterGroup, controller *controllers.BiographyController) {
	bioRoutes := routerGroup.Group("/biographies")
	{
		bioRoutes.GET("", controller.GetAll)
		bioRoutes.GET("/:id", controller.GetByID)
		bioRoutes.POST("", controller.Create)
		bioRoutes.PUT("/:id", controller.Update)
		bioRoutes.DELETE("/:id", controller.Delete)
	}
}
