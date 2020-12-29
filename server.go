package main

import (
	controller "Inexpediency/simple-gin-rest/contoller"
	"Inexpediency/simple-gin-rest/middleware"
	"Inexpediency/simple-gin-rest/service"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()
	server.Use(
		gin.Recovery(),
		gin.Logger(),
		middleware.BasicAuth(),
	)

	server.POST("/video", videoController.Save)
	server.GET("/videos", videoController.FindAll)

	err := server.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
