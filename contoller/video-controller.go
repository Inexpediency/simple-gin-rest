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

type controller struct {
	service service.VideoService
}

// New returns new video controller
func New(service service.VideoService) VideoController {
	return controller{
		service: service,
	}
}

// Save saves video
func (controller controller) Save(ctx *gin.Context) {
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
func (controller controller) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, controller.service.FindAll())
}
