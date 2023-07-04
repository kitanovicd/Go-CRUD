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

	r.POST("/signup", controllers.SignUp)
	r.PUT("/user/:id", controllers.UpdateUser)
	r.GET("/user/:id", controllers.GetUser)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:company_id", controllers.GetUsersByCompany)
	r.DELETE("/user/:id", controllers.DeleteUser)

	r.Run()
}
