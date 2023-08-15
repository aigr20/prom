package api

import (
	"aigr20/prom/database"
	"aigr20/prom/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api *API) GetIssueHandler(ctx *gin.Context) {
	var issueId int
	if id, ok := ctx.Params.Get("issueId"); ok {
		var err error
		issueId, err = strconv.Atoi(id)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
	} else {
		log.Println("Missing issueId")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	issue, err := api.IssueRepo.GetOne(issueId)
	if err != nil {
		var status int
		switch err {
		case database.ErrIssueNotFound:
			status = http.StatusNotFound
		default:
			status = http.StatusBadRequest
		}
		ctx.AbortWithStatus(status)
		return
	}

	ctx.JSON(http.StatusOK, ResponseData{"data": issue})
}

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

func (api *API) AddIssueTagsHandler(ctx *gin.Context) {
	var body models.IssueTagsBody
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	issue, err := api.IssueRepo.GetOne(body.IssueID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	for i := range body.Tags {
		if _, ok := api.TagRepo.FindTagInProject(body.Tags[i], issue.ProjectID); !ok {
			fmt.Printf("Tag %v is not part of project %v\n", body.Tags[i], issue.ProjectID)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}

	err = api.IssueRepo.AddTags(issue.ID, body.Tags)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (api *API) RemoveIssueTagsHandler(ctx *gin.Context) {
	var body models.IssueTagsBody
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = api.IssueRepo.RemoveTags(body.IssueID, body.Tags)
	if err != nil {
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
