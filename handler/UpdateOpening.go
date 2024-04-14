package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hcncc/gopportunities-api/schemas"
)

func UpdateOpeningHandler(context *gin.Context) {
	request := UpdateRequestOpening{}
	context.BindJSON(&request)

	if err := request.ValidateFields(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())

		sendError(context, http.StatusBadRequest, err.Error())

		return
	}

	id := context.Query("id")

	if id == "" {
		sendError(context, http.StatusBadRequest, errorParameterIsRequired("id", "QueryParameter").Error())

		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(context, http.StatusNotFound, "Opening not found")
		return
	}

	// Mesclando os dados
	opening.Role = request.Role
	opening.Company = request.Company
	opening.Link = request.Link
	opening.Location = request.Location
	opening.Salary = request.Salary
	opening.Remote = *request.Remote

	if err := db.Save(&opening).Error; err != nil {
		sendError(context, http.StatusInternalServerError, "Error in Update Opening")

		return
	}

	sendSuccess(context, "Update", opening)
}
