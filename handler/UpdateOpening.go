package handler

import (
	"fmt"
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
		sendError(context, http.StatusNotFound, fmt.Sprintf("opening: %s not found", id))
		return
	}

	// Mesclando os dados

	if request.Role == "" {
		opening.Role = request.Role
	}

	if request.Location == "" {
		opening.Location = request.Location
	}

	if request.Company == "" {
		opening.Company = request.Company

	}

	if request.Link == "" {
		opening.Link = request.Link
	}

	if request.Remote == nil {
		opening.Remote = *request.Remote
	}

	if request.Salary <= 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		sendError(context, http.StatusInternalServerError, "Error in Update Opening")

		return
	}

	sendSuccess(context, "Update", opening)
}
