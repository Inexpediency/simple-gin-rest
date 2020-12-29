package controller

import (
	"Inexpediency/simple-gin-rest/entity"
	"Inexpediency/simple-gin-rest/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// VideoController implementation
type VideoController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
}

type videoController struct {
	service service.VideoService
}

// NewVideoController returns new video videoController
func NewVideoController() VideoController {
	return videoController{
		service: service.NewVideoService(),
	}
}

// Save saves video
func (controller videoController) Save(ctx *gin.Context) {
	var video entity.Video
	err := ctx.BindJSON(&video)
	log.Println(video)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	controller.service.Save(video)
	ctx.JSON(http.StatusOK, video)
}

// FindAll returns all added videos
func (controller videoController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, controller.service.FindAll())
}
