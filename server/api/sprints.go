package api

import (
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

	sprint, err := api.SprintRepo.GetSprint(sprintId)
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
