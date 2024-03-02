package models

type RealEstate struct {
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
