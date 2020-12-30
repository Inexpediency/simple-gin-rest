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
// CreateVideo godoc
// @Security bearerAuth
// @Summary Create new videos
// @Description Create a new video
// @Tags videos,create
// @Accept  json
// @Produce  json
// @Router /videos [post]
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
// GetVideos godoc
// @Security bearerAuth
// @Summary List existing videos
// @Description Get all the existing videos
// @Tags videos,list
// @Accept  json
// @Produce  json
// @Router /videos [get]
func (controller videoController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, controller.service.FindAll())
}

// Update updates video in db
// UpdateVideo godoc
// @Security bearerAuth
// @Summary Update videos
// @Description Update a single video
// @Security bearerAuth
// @Tags videos
// @Accept  json
// @Produce  json
// @Param  id path int true "Video ID"
// @Router /videos/{id} [PATCH]
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
// DeleteVideo godoc
// @Security bearerAuth
// @Summary Remove videos
// @Description Delete a single video
// @Security bearerAuth
// @Tags videos
// @Accept  json
// @Produce  json
// @Param  id path int true "Video ID"
// @Router /videos/{id} [delete]
func (controller videoController) Delete(ctx *gin.Context) {
	var video entity.Video

	id, err := strconv.ParseUint(ctx.Param("id"),10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	video.ID = id

	controller.service.Delete(video)

	ctx.JSON(http.StatusOK, gin.H{})
}
