package api

import (
	"aigr20/prom/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *API) CreateIssueHandler(ctx *gin.Context) {
	var body models.IssueCreateForm
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		log.Println(err)
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

func (api *API) UpdateIssueStatusHandler(ctx *gin.Context) {
	var body models.UpdateIssueStatusBody
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	statusID, err := api.StatusRepo.GetIDByName(body.NewStatus)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	_, err = api.IssueRepo.UpdateIssue(body.IssueID, []string{"issue_status"}, []any{statusID})
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	ctx.JSON(http.StatusNoContent, nil)
}
