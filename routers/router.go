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
			userGroup.GET("/hello", controllers.TestHandler)
		}
	}
}
