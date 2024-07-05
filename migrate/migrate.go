package main

import (
	"go-http/initializers"
	"go-http/models"
)

func init() {
	initializers.LoadEnvVeriables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
