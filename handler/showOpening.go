package handler

import (
	"net/http"

	"github.com/eduardylopes/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusNotFound, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	OpeningResponse := schemas.OpeningResponse{
		ID:       opening.ID,
		Role:     opening.Role,
		Location: opening.Location,
		Company:  opening.Company,
		Remote:   opening.Remote,
		Link:     opening.Link,
		Salary:   opening.Salary,
		CreateAt: opening.CreatedAt,
		UpdateAt: opening.UpdatedAt,
	}

	sendSuccess(ctx, "show-opening", OpeningResponse)
}
