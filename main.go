package main

import (
	"github.com/gin-gonic/gin"
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
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}