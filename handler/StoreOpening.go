package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hcncc/gopportunities-api/schemas"
)

// @BasePath /api
// @Summary Create Opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request Body"
// @Success 201 {object} StoreOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func StoreOpeningHandler(context *gin.Context) {
	request := CreateOpeningRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())

		sendError(context, http.StatusBadRequest, err.Error())
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error in creating opening: %v", err.Error())
		sendError(context, http.StatusInternalServerError, "Error creating opening in Database because offline DB")
		return
	}

	sendSuccess(context, "Store", opening)
}
