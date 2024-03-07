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

	if err != nil {
		log.Fatal("Wystąpił błąd podczas migracji bazy danych")
	}
}
