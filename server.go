package main

import (
	controller "Inexpediency/simple-gin-rest/contoller"
	"Inexpediency/simple-gin-rest/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.POST("/video", func(ctx *gin.Context) {
		res, err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, res)
	})

	server.GET("/videos", func(ctx *gin.Context) {
		res, err := videoController.FindAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, res)
	})

	err := server.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
