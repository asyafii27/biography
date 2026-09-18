package routes

import (
	"biography-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterSkillRoutes(routerGroup *gin.RouterGroup, controller *controllers.SkillController) {
	routes := routerGroup.Group("/skills")
	{
		routes.GET("", controller.GetAll)
		routes.GET("/:id", controller.GetByID)
		routes.POST("", controller.Create)
		routes.PUT("/:id", controller.Update)
		routes.DELETE("/:id", controller.Delete)
	}
}
