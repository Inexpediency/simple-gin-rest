package main

import (
	controller "Inexpediency/simple-gin-rest/contoller"
	"Inexpediency/simple-gin-rest/docs"
	"Inexpediency/simple-gin-rest/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
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
	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Ythosa - Video API"
	docs.SwaggerInfo.Description = "Ythosa - Youtube Gin Video API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"https"}

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

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := server.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
