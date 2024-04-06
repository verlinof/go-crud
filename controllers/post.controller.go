package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/verlinof/go-crud/initializers"
	"github.com/verlinof/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	//Get data off req body

	//Create a post
	post := models.Post{
		Title: c.PostForm("title"),
		Body:  c.PostForm("body"),
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(500)
		return
	}
	//Return it
	c.JSON(200, gin.H{
		"message": "Post created",
		"post": post,
	})
}