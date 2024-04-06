package main

import (
	"github.com/verlinof/go-crud/initializers"
	"github.com/verlinof/go-crud/models"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
}

func main() {
	// Migrate the schema
	initializers.DB.AutoMigrate(&models.Post{})
}