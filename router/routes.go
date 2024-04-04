package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	routerApi := router.Group("/api")

	routerApi.GET("/opening", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "GET Opening Route",
		})
	})

	routerApi.POST("/opening", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "POST Opening Route",
		})
	})

	routerApi.PUT("/opening", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "PUT Opening Route",
		})
	})

	routerApi.DELETE("/opening", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "DELETE Opening Route",
		})
	})

	routerApi.GET("/openings", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "GET Openings Route",
		})
	})
}
