package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kitanovicd/Go-CRUD/Go-CRUD/controllers"
	"github.com/kitanovicd/Go-CRUD/Go-CRUD/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {

	r := gin.Default()

	r.POST("/company", controllers.CreateCompany)
	r.PUT("/company/:id", controllers.UpdateCompany)
	r.GET("/company/:id", controllers.GetCompany)
	r.GET("/companies", controllers.GetCompanies)
	r.DELETE("/company/:id", controllers.DeleteCompany)

	r.Run()
}
