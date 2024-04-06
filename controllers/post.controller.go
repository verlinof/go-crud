package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/verlinof/go-crud/initializers"
	"github.com/verlinof/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	//Get data off req body
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	//Create a post
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Failed to create post",
			"error":   result.Error.Error(),
		})
		return
	}
	//Return it
	c.JSON(200, gin.H{
		"message": "Post created",
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//Get the posts
	// List for Posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Return response
	c.JSON(200, gin.H{
		"message" : "success",
		"posts": posts,
	})
}