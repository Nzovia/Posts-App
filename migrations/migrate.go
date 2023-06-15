package main

import (
	"personalprojects/Backend/Go-Programming/go-crud/initializers"
	"personalprojects/Backend/Go-Programming/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectionToDatabase()
}

// for auto migration
func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}