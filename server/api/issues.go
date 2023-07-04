package api

import (
	"aigr20/prom/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *API) CreateIssueHandler(ctx *gin.Context) {
	var body models.IssueCreateForm
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	issue, err := api.IssueRepo.CreateIssue(body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, ResponseData{"data": issue})
}
