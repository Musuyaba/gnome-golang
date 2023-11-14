package routers

import (
	"github.com/Musuyaba/gnome-golang/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func ApiRoutes(superRouter *gin.Engine) {
	apiGroup := superRouter.Group("/api")
	{
		userGroup := apiGroup.Group("/user")
		{
			userGroup.POST("/", controllers.TestHandler)
			userGroup.DELETE("/", controllers.TestHandler)
			userGroup.GET("/", controllers.TestHandler)
		}
	}
}

func PublicRoutes(superRouter *gin.Engine) {
	publicApiGroup := superRouter.Group("/")
	{
		publicApiGroup.GET("/", controllers.TestHandler)
		// userGroup := publicApiGroup.Group("/user")
		// {
		// 	userGroup.GET("/hello", controllers.TestHandler)
		// }
	}
}
