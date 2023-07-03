package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kitanovicd/Go-CRUD/Go-CRUD/initializers"
	"github.com/kitanovicd/Go-CRUD/Go-CRUD/models"
)

type CompanyBody struct {
	Name          string
	ContactPerson string
}

func CreateCompany(c *gin.Context) {
	var body CompanyBody
	c.Bind(&body)

	company := models.Company{
		Name:          body.Name,
		ContactPerson: body.ContactPerson,
	}
	result := initializers.DB.Create(&company)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"company": company,
	})
}

func GetCompanies(c *gin.Context) {
	var companies []models.Company
	result := initializers.DB.Find(&companies)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"companies": companies,
	})
}

func GetCompany(c *gin.Context) {
	var company models.Company
	result := initializers.DB.First(&company, c.Param("id"))

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"company": company,
	})
}

func UpdateCompany(c *gin.Context) {
	var company models.Company
	result := initializers.DB.First(&company, c.Param("id"))

	if result.Error != nil {
		c.Status(400)
		return
	}

	var body CompanyBody
	c.Bind(&body)

	initializers.DB.Model(&company).Updates(models.Company{
		Name:          body.Name,
		ContactPerson: body.ContactPerson,
	})

	c.JSON(200, gin.H{
		"company": company,
	})
}

func DeleteCompany(c *gin.Context) {
	initializers.DB.Delete(&models.Company{}, c.Param("id"))
	c.Status(200)
}
