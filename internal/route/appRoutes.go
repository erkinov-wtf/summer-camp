package route

import (
	"camp-summer/internal/api/dependencyIndejection"
	"github.com/gin-gonic/gin"
)

func SetupAppRoutes(router *gin.RouterGroup, container *dependencyIndejection.Container) {
	appRoutes := router.Group("/")
	{
		appRoutes.POST("/create/", container.AppController.Create)
		appRoutes.GET("/:id", container.AppController.GetById)
		appRoutes.PATCH("/:id", container.AppController.Patch)
	}
}
