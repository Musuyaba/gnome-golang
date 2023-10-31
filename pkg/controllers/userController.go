package controllers

import "github.com/gin-gonic/gin"

func TestHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "testhandler",
	})
}
