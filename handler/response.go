package handler

import (
	"github.com/gin-gonic/gin"
)

func sendError(context *gin.Context, code int, message string) {

	context.Header("Content-type", "Application/json")

	context.JSON(code, gin.H{
		"message":    message,
		"statusCode": code,
	})
}
