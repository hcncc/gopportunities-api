package router

import "github.com/gin-gonic/gin"

func InitializeServer() {
	app := gin.Default()

	// Initialize routes application
	initializeRoutes(app)

	app.Run(":8080")
}
