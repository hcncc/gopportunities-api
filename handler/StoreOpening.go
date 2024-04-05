package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StoreOpeningHandler(context *gin.Context) {
	request := CreateOpeningRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		
		sendError(context, http.StatusBadRequest, err.Error())
	}

	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("Error in creating opening: %v", err.Error())

		return
	}
}
