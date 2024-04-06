package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/verlinof/go-crud/initializers"
	"github.com/verlinof/go-crud/models"
)

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

func PostsShow(c *gin.Context) {
	//Get Post by ID
	var post models.Post
	//c.Param digunakan untuk mengambil parameter dari URL
	initializers.DB.First(&post, c.Param("id"))

	//Return response
	c.JSON(200, gin.H{
		"message" : "success",
		"post": post,
	})
}

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

func PostUpdate(c *gin.Context) {
	//Get data from BODY
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)
	
	//Find and Update
	var post models.Post
	initializers.DB.First(&post, c.Param("id"))
	result := initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title, 
		Body: body.Body,
	})

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Failed to create post",
			"error":   result.Error.Error(),
		})
		return
	}

	//Return response
	c.JSON(200, gin.H{
		"message" : "success",
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	//Get post by ID
	var post models.Post
	initializers.DB.First(&post, c.Param("id"))
	//Delete the data
	result := initializers.DB.Delete(&post, c.Param("id"))

	if(result.Error != nil) {
		c.JSON(500, gin.H{
			"message": "Failed to delete post",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Post deleted",
		"post": post,
	})
}