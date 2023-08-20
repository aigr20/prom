package api

import (
	"aigr20/prom/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api *API) GetProjectsHandler(ctx *gin.Context) {
	projects, err := api.ProjectRepo.GetAll()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, ErrGetProjects)
		return
	}

	ctx.JSON(200, ResponseData{"data": projects})
}

func (api *API) GetProjectHandler(ctx *gin.Context) {
	projectId, err := strconv.Atoi(ctx.Param("projectId"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, ErrBadRequest)
		return
	}

	project, err := api.ProjectRepo.GetOne(projectId)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, ErrGetProjects)
		return
	}

	ctx.JSON(200, ResponseData{"data": project})
}

func (api *API) GetProjectIssuesHandler(ctx *gin.Context) {
	projectId, err := strconv.Atoi(ctx.Param("projectId"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, ErrBadRequest)
		return
	}

	issues, err := api.IssueRepo.GetIssuesFromProject(projectId)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, ResponseData{"data": issues})
}

func (api *API) CreateProjectHandler(ctx *gin.Context) {
	var body models.ProjectCreateForm
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, ErrBadRequest)
		return
	}

	project, err := api.ProjectRepo.CreateProject(body)
	if err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, ResponseData{"data": project})
}

func (api *API) GetTagCountsHandler(ctx *gin.Context) {
	projectId, err := strconv.Atoi(ctx.Param("projectId"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	counts, err := api.ProjectRepo.GetTagIssueCounts(projectId)
	if err != nil {
		log.Println("Failed to count tags for project ", projectId)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, ResponseData{"data": counts})
}
