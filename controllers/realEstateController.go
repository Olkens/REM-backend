package controllers

import (
	"github.com/REM/initializers"
	"github.com/REM/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveRealEstate(ctx *gin.Context) {
	var realEstate models.RealEstate
	err := ctx.ShouldBindJSON(&realEstate)
	result := initializers.DB.Create(&realEstate)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Błąd podczas dodawania nieruchomości"})
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Dodano nową nieruchomość", "realEstate": result})
}
