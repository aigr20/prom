package api

import (
	"aigr20/prom/database"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type ResponseData = map[string]interface{}

type API struct {
	Router      *gin.Engine
	ProjectRepo database.ProjectRepository
}

func NewAPI(db *sql.DB) *API {
	api := &API{
		Router:      gin.Default(),
		ProjectRepo: *database.NewProjectRepository(db),
	}
	api.Routes()
	return api
}

func (api *API) Serve() error {
	return api.Router.Run(":8080")
}

func (api *API) Routes() {
	projectsGroup := api.Router.Group("/projects")
	{
		projectsGroup.GET("/all", api.GetProjectsHandler)
	}
}
