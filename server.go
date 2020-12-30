package main

import (
	controller "Inexpediency/simple-gin-rest/contoller"
	"Inexpediency/simple-gin-rest/middleware"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

var (
	videoController = controller.NewVideoController()
	loginController = controller.NewLoginController()
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
		//middleware.BasicAuth(),
	)

	server.POST("/login", loginController.Login)

	apiRoutes := server.Group("/api", middleware.AuthorizeJWT())
	{
		apiRoutes.POST("/video", videoController.Save)
		apiRoutes.PATCH("/video/:id", videoController.Update)
		apiRoutes.DELETE("/video/:id", videoController.Delete)
		apiRoutes.GET("/videos", videoController.FindAll)
	}

	err := server.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
