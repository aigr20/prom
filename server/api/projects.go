package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ErrGetProjects = errors.New("failed retrieving projects")

func (api *API) GetProjectsHandler(ctx *gin.Context) {
	projects, err := api.ProjectRepo.GetAll()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, ErrGetProjects)
		return
	}

	ctx.JSON(200, ResponseData{"data": projects})
}
