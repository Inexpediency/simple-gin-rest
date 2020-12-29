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

	server.POST("/add-video", func(ctx *gin.Context) {
		res, err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, struct {
				error error
			}{
				error: err,
			})
		}

		ctx.JSON(http.StatusOK, res)
	})

	err := server.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
