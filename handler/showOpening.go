package handler

import (
	"net/http"

	"github.com/eduardylopes/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

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
