package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(context *gin.Context, code int, message string) {

	context.Header("Content-type", "application/json")

	context.JSON(code, gin.H{
		"message":    message,
		"statusCode": code,
	})
}

func sendSuccess(context *gin.Context, operation string, data ...interface{}) {
	context.Header("Content-type", "application/json")

	if operation == "Store" {
		context.JSON(http.StatusCreated, gin.H{
			"message": "Created success",
			"data":    data,
		})
	}

	if operation == "Update" {
		context.JSON(http.StatusOK, gin.H{
			"message": "Update success",
			"data":    data,
		})
	}

	if operation == "Delete" {
		context.JSON(http.StatusOK, gin.H{
			"message": "Deleted success",
			"data":    data,
		})
	}

	if operation == "All" {
		context.JSON(http.StatusOK, gin.H{
			"data": data[0],
		})
	}

	if operation == "Unique" {
		context.JSON(http.StatusFound, gin.H{
			"data": data,
		})
	}
}
