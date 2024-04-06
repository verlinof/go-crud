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
	r.POST("/posts", controllers.PostsCreate)

	r.Run()
}