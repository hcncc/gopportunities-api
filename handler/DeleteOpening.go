package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hcncc/gopportunities-api/schemas"
)

func DeleteOpeningHandler(context *gin.Context) {
	id := context.Query("id")

	if id == "" {
		sendError(context, http.StatusBadRequest, errorParameterIsRequired("id", "QueryParameter").Error())

		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(context, http.StatusNotFound, fmt.Sprintf("Opening id: %s not found", id))

		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(context, http.StatusInternalServerError, fmt.Sprintf("Error deleting opening with id: %s", id))

		return
	}

	sendSuccess(context, "Delete", opening)
}
