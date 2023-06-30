package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ErrGetProjects = errors.New("failed retrieving projects")
var ErrBadRequest = errors.New("bad request")

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
