package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kitanovicd/Go-CRUD/Go-CRUD/initializers"
	"github.com/kitanovicd/Go-CRUD/Go-CRUD/models"
)

type UserBody struct {
	Name      string
	Surname   string
	Email     string
	Phone     string
	CompanyID uint
}

func CreateUser(c *gin.Context) {
	var body UserBody
	c.Bind(&body)

	user := models.User{
		Name:      body.Name,
		Surname:   body.Surname,
		Email:     body.Email,
		Phone:     body.Phone,
		CompanyID: body.CompanyID,
	}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	result := initializers.DB.Find(&users)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}

func GetUsersByCompany(c *gin.Context) {
	var users []models.User
	result := initializers.DB.Where("company_id = ?", c.Param("company_id")).Find(&users)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}

func GetUser(c *gin.Context) {
	var user models.User
	result := initializers.DB.First(&user, c.Param("id"))

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	result := initializers.DB.First(&user, c.Param("id"))

	if result.Error != nil {
		c.Status(400)
		return
	}

	var body UserBody
	c.Bind(&body)

	initializers.DB.Model(&user).Updates(models.User{
		Name:      body.Name,
		Surname:   body.Surname,
		Email:     body.Email,
		Phone:     body.Phone,
		CompanyID: user.CompanyID,
	})

	c.JSON(200, gin.H{
		"user": user,
	})
}

func DeleteUser(c *gin.Context) {
	initializers.DB.Delete(&models.User{}, c.Param("id"))
	c.Status(200)
}
