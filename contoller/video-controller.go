package controller

import (
	"Inexpediency/simple-gin-rest/entity"
	"Inexpediency/simple-gin-rest/service"
	"github.com/gin-gonic/gin"
)

// VideoController implementation
type VideoController interface {
	FindAll() ([]entity.Video, error)
	Save(ctx *gin.Context) (entity.Video, error)
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
func (controller controller) Save(ctx *gin.Context) (entity.Video, error) {
	var video entity.Video
	err := ctx.BindJSON(&video)
	if err != nil {
		return entity.Video{}, err
	}

	controller.service.Save(video)

	return video, nil
}

// FindAll returns all added videos
func (controller controller) FindAll() ([]entity.Video, error) {
	return controller.service.FindAll(), nil
}
