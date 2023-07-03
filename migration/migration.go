package main

import (
	"github.com/kitanovicd/Go-CRUD/Go-CRUD/initializers"
	"github.com/kitanovicd/Go-CRUD/Go-CRUD/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Company{})

}
