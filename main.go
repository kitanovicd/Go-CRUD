package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kitanovicd/Go-CRUD/Go-CRUD/controllers"
	"github.com/kitanovicd/Go-CRUD/Go-CRUD/initializers"
	"github.com/kitanovicd/Go-CRUD/Go-CRUD/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {

	r := gin.Default()

	r.POST("/company", middleware.RequireAuth, controllers.CreateCompany)
	r.PUT("/company/:id", middleware.RequireAuth, controllers.UpdateCompany)
	r.GET("/company/:id", middleware.RequireAuth, controllers.GetCompany)
	r.GET("/companies", middleware.RequireAuth, controllers.GetCompanies)
	r.DELETE("/company/:id", middleware.RequireAuth, controllers.DeleteCompany)

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.PUT("/user/:id", middleware.RequireAuth, controllers.UpdateUser)
	r.GET("/user/:id", middleware.RequireAuth, controllers.GetUser)
	r.GET("/users", middleware.RequireAuth, controllers.GetUsers)
	r.GET("/users/:company_id", middleware.RequireAuth, controllers.GetUsersByCompany)
	r.DELETE("/user/:id", middleware.RequireAuth, controllers.DeleteUser)

	r.Run()
}
