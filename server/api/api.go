package api

import (
	"aigr20/prom/database"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData = map[string]interface{}

type API struct {
	Router      *gin.Engine
	ProjectRepo database.ProjectRepository
	IssueRepo   database.IssueRepository
	StatusRepo  database.StatusRepository
	TagRepo     database.TagRepository
}

func NewAPI(db *sql.DB) *API {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	api := &API{
		Router:      router,
		ProjectRepo: *database.NewProjectRepository(db),
		IssueRepo:   *database.NewIssueRepository(db),
		StatusRepo:  *database.NewStatusRepository(db),
		TagRepo:     *database.NewTagRepository(db),
	}
	api.Routes()
	return api
}

func CorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Next()
	}
}

func (api *API) Serve() error {
	return api.Router.Run(":8080")
}

func (api *API) Routes() {
	api.Router.Use(CorsMiddleware())
	projectsGroup := api.Router.Group("/projects")
	{
		projectsGroup.GET("/all", api.GetProjectsHandler)
		projectsGroup.GET("/:projectId", api.GetProjectHandler)
		projectsGroup.GET("/:projectId/issues", api.GetProjectIssuesHandler)
		projectsGroup.POST("/create", api.CreateProjectHandler)
	}
	issuesGroup := api.Router.Group("/issues")
	{
		issuesGroup.GET("/:issueId", api.GetIssueHandler)
		issuesGroup.POST("/create", api.CreateIssueHandler)
		issuesGroup.PATCH("/status", api.UpdateIssueStatusHandler)
		issuesGroup.PATCH("/update", api.UpdateIssueHandler)
		issuesGroup.PATCH("/tags", api.AddIssueTagsHandler)
		issuesGroup.DELETE("/tags", api.RemoveIssueTagsHandler)
	}
	tagsGroup := api.Router.Group("/tags")
	{
		tagsGroup.POST("/create", api.CreateTagHandler)
	}
}
