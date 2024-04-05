package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hcncc/gopportunities-api/handler"
)

func initializeRoutes(router *gin.Engine) {

	//Initialize Handler
	handler.InitializeHandler()

	routerApi := router.Group("/api")

	routerApi.GET("/opening", handler.ShowOpeningHandler)

	routerApi.POST("/opening", handler.StoreOpeningHandler)

	routerApi.PUT("/opening", handler.UpdateOpeningHandler)

	routerApi.DELETE("/opening", handler.DeleteOpeningHandler)

	routerApi.GET("/openings", handler.IndexOpeningHandler)
}
