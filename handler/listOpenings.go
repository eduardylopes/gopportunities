package handler

import (
	"net/http"

	"github.com/eduardylopes/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
	}

	openingsRes := []schemas.OpeningResponse{}

	for _, opening := range openings {
		openingsRes = append(openingsRes, schemas.OpeningResponse{
			ID:       opening.ID,
			Role:     opening.Role,
			Location: opening.Location,
			Company:  opening.Company,
			Remote:   opening.Remote,
			Link:     opening.Link,
			Salary:   opening.Salary,
			CreateAt: opening.CreatedAt,
			UpdateAt: opening.UpdatedAt,
		})
	}

	sendSuccess(ctx, "list-openings", openingsRes)
}
