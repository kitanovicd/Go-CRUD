package controllers

import (
	"net/http"

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
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid body",
		})
		return
	}

	company := models.Company{
		Name:          body.Name,
		ContactPerson: body.ContactPerson,
	}
	result := initializers.DB.Create(&company)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
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

	c.JSON(http.StatusOK, gin.H{
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

	c.JSON(http.StatusOK, gin.H{
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
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid body",
		})
		return
	}

	result = initializers.DB.Model(&company).Updates(models.Company{
		Name:          body.Name,
		ContactPerson: body.ContactPerson,
	})
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"company": company,
	})
}

func DeleteCompany(c *gin.Context) {
	result := initializers.DB.Delete(&models.Company{}, c.Param("id"))
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.Status(http.StatusOK)
}
