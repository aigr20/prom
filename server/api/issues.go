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
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (api *API) UpdateIssueHandler(ctx *gin.Context) {
	var body models.UpdateIssueBody
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fields := make([]string, 0)
	values := make([]any, 0)
	for field, value := range body.Updates {
		switch field {
		case "title", "description":
			fields = append(fields, "issue_"+field)
			values = append(values, value)
		default:
			fields = append(fields, field)
			values = append(values, value)
		}
	}

	_, err = api.IssueRepo.UpdateIssue(body.IssueID, fields, values)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.Status(http.StatusNoContent)
}
