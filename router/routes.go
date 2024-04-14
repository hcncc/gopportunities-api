package router

import (
	"github.com/gopportunities-api/docs"
	"github.com/hcncc/gopportunities-api/handler"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {

	//Initialize Handler
	handler.InitializeHandler()

	basePath := "/api"
	docs.SwaggerInfo.BasePath = basePath

	routerApi := router.Group(basePath)

	routerApi.GET("/opening", handler.ShowOpeningHandler)

	routerApi.POST("/opening", handler.StoreOpeningHandler)

	routerApi.PUT("/opening", handler.UpdateOpeningHandler)

	routerApi.DELETE("/opening", handler.DeleteOpeningHandler)

	routerApi.GET("/openings", handler.IndexOpeningHandler)
	routerApi.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
