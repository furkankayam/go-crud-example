package controllers

import (
	"github.com/gin-gonic/gin"
	"go-http/initializers"
	"go-http/models"
)

// PostsCreate godoc
// @Summary Create a new post
// @Description Creates a new post in the database.
// @Tags posts
// @Accept json
// @Produce json
// @Param body body models.Post true "Post data"
// @Success 200 {object} models.Post "The created post"
// @Router /posts [post]
func PostsCreate(c *gin.Context) {

	// Get data off req body
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the Post
	var post models.Post

	initializers.DB.Find(&post, id)

	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	//Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)

}
