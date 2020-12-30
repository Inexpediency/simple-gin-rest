package controller

import (
	"Inexpediency/simple-gin-rest/entity"
	"Inexpediency/simple-gin-rest/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// VideoController implementation
type VideoController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
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

	err := ctx.ShouldBindJSON(&video)
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

// Update updates video in db
func (controller videoController) Update(ctx *gin.Context) {
	var video entity.Video

	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	id, err := strconv.ParseUint(ctx.Param("id"),10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	video.ID = id

	controller.service.Update(video)

	ctx.JSON(http.StatusOK, gin.H{})
}

// Delete deletes video from db
func (controller videoController) Delete(ctx *gin.Context) {
	var video entity.Video

	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	id, err := strconv.ParseUint(ctx.Param("id"),10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	video.ID = id

	controller.service.Delete(video)

	ctx.JSON(http.StatusOK, gin.H{})
}
