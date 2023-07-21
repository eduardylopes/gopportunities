package handler

import (
	"net/http"
	"time"

	"github.com/eduardylopes/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
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
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	openingRes := schemas.OpeningResponse{
		ID:        opening.ID,
		Role:      opening.Role,
		Location:  opening.Location,
		Company:   opening.Company,
		Remote:    opening.Remote,
		Link:      opening.Link,
		Salary:    opening.Salary,
		CreateAt:  opening.CreatedAt,
		UpdateAt:  opening.UpdatedAt,
		DeletedAt: time.Time{},
	}

	sendSuccess(ctx, "create-opening", openingRes)
}
