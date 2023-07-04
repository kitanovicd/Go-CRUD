package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string
	Surname   string
	Username  string
	Password  string
	Email     string `gorm:"unique"`
	Phone     string
	CompanyID uint
}
