package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StoreOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "POST Opening Route",
	})
}
