package main

import (
	controller "Inexpediency/simple-gin-rest/contoller"
	"Inexpediency/simple-gin-rest/service"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.POST("/video", videoController.Save)
	server.GET("/videos", videoController.FindAll)

	err := server.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
