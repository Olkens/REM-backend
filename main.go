package main

import (
	"github.com/REM/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	server.GET("/real-estates", getAllRealEstates)
	server.POST("/real-estates", createRealEstate)

	err := server.Run(":8080")
	if err != nil {
		return
	}
}

func getAllRealEstates(context *gin.Context) {
	realEstates := models.GetAllRealEstate()
	context.JSON(http.StatusOK, realEstates)
}

func createRealEstate(context *gin.Context) {
	var realEstate models.RealEstate
	err := context.ShouldBindJSON(&realEstate)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Błąd podczas dodawania nieruchomości"})
	}
	realEstate.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Dodano nową nieruchomość", "realEstate": realEstate})
}
