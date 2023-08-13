package api

import (
	"aigr20/prom/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *API) CreateTagHandler(ctx *gin.Context) {
	var body models.CreateTagBody
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	tagId, err := api.TagRepo.CreateTag(body)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	tag, ok := api.TagRepo.FindTagByID(tagId)
	if !ok {
		log.Printf("Failed to find newly created tag with id %v\n", tagId)
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusCreated, ResponseData{"data": tag})
}
