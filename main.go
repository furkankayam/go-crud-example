package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-http/controllers"
	_ "go-http/docs"
	"go-http/initializers"
)

func init() {
	initializers.LoadEnvVeriables()
	initializers.ConnectToDB()
}

// @title Post Swagger
// @version 1.0
// @description this is sample
// @termsOfService

// @contact.name furkan
// @contact.url
// @contact.email

// @license.name
// @license.url

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	r := gin.Default()

	// Swagger
	url := ginSwagger.URL("http://localhost:3000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.GET("/posts", controllers.PostIndex)
	r.GET("/posts/:id", controllers.PostShow)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
	//r.Run("8080")
}
