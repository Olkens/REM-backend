package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        int
	FirstName string `binding:"required"`
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
	Enabled   bool
	Locked    bool
}

var user = []User{}

func GetAllUsers() []User {
	return user
}
