package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hcncc/gopportunities-api/schemas"
)

func IndexOpeningHandler(context *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(context, http.StatusInternalServerError, "error listening openings")
		return
	}

	sendSuccess(context, "All", openings)
}
