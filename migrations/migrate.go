package main

import (
	"github.com/REM/initializers"
	"github.com/REM/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.RealEstate{})
	userErr := initializers.DB.AutoMigrate(&models.User{})

	if err != nil || userErr != nil {
		log.Fatal("Wystąpił błąd podczas migracji bazy danych")
	}
}
