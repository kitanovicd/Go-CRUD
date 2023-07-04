package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kitanovicd/Go-CRUD/Go-CRUD/initializers"
	"github.com/kitanovicd/Go-CRUD/Go-CRUD/models"
	"golang.org/x/crypto/bcrypt"
)

type UserBody struct {
	Name      string
	Surname   string
	Username  string
	Password  string
	Email     string
	Phone     string
	CompanyID uint
}

func SignUp(c *gin.Context) {
	var body UserBody
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to hash password",
		})
		return
	}

	user := models.User{
		Name:      body.Name,
		Surname:   body.Surname,
		Username:  body.Username,
		Password:  string(hash),
		Email:     body.Email,
		Phone:     body.Phone,
		CompanyID: body.CompanyID,
	}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func SignIn(c *gin.Context) {
	var body struct {
		Username string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid body",
		})
		return
	}

	var user models.User
	result := initializers.DB.Where("username = ?", body.Username).First(&user)

	if result.Error != nil || user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid password",
		})
		return
	}

}

func GetUsers(c *gin.Context) {
	var users []models.User
	result := initializers.DB.Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get users",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetUsersByCompany(c *gin.Context) {
	var users []models.User
	result := initializers.DB.Where("company_id = ?", c.Param("company_id")).Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get users",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetUser(c *gin.Context) {
	var user models.User
	result := initializers.DB.First(&user, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	result := initializers.DB.First(&user, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get user",
		})
		return
	}

	var body UserBody
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid body",
		})
		return
	}

	result = initializers.DB.Model(&user).Updates(models.User{
		Name:      body.Name,
		Surname:   body.Surname,
		Username:  body.Username,
		Password:  body.Password,
		Email:     body.Email,
		Phone:     body.Phone,
		CompanyID: user.CompanyID,
	})
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func DeleteUser(c *gin.Context) {
	result := initializers.DB.Delete(&models.User{}, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete user",
		})
		return
	}

	c.Status(http.StatusOK)
}
