package main

import (
	"github.com/gin-gonic/gin"
	"github.com/verlinof/go-crud/controllers"
	"github.com/verlinof/go-crud/initializers"
)

//Init akan dijalankan sebelum main
func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	//Routes
	//Post
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)
	r.POST("/posts", controllers.PostsCreate)

	r.Run()
}