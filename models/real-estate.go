package models

import "gorm.io/gorm"

type RealEstate struct {
	gorm.Model
	ID   int
	Name string `binding:"required"`
}

var realEstates = []RealEstate{}

func (re RealEstate) Save() {
	realEstates = append(realEstates, re)
}

func GetAllRealEstate() []RealEstate {
	return realEstates
}
