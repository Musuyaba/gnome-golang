package routers

import "github.com/gin-gonic/gin"

func ApiRoutes(superRouter *gin.Engine) {
	apiGroup := superRouter.Group("/api")
	{
		userGroup := apiGroup.Group("/user")
		{
			userGroup.GET("/hello", func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{
					"message": "mantap",
				})
			})
		}
	}
}
