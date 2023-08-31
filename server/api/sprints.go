package api

import (
	"aigr20/prom/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api *API) GetSprintHandler(ctx *gin.Context) {
	sprintId, err := strconv.Atoi(ctx.Param("sprintId"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	sprint, err := api.SprintRepo.GetSprintWithIssues(sprintId)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, ResponseData{"data": sprint})
}

func (api *API) GetSprintIssuesHandler(ctx *gin.Context) {
	sprintId, err := strconv.Atoi(ctx.Param("sprintId"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	issues, err := api.SprintRepo.GetIssues(sprintId)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, ResponseData{"data": issues})
}

func (api *API) CreateSprintHandler(ctx *gin.Context) {
	var body models.CreateSprintBody
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, err := api.SprintRepo.CreateSprint(body)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	created, err := api.SprintRepo.GetSprint(id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusCreated, ResponseData{"data": created})
}
