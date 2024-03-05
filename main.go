package main

import (
	"github.com/REM/controllers"
	"github.com/REM/initializers"
	"github.com/REM/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	server := gin.Default()

	server.GET("/real-estates", getAllRealEstates)
	server.POST("/real-estates", controllers.SaveRealEstate)

	err := server.Run()
	if err != nil {
		return
	}
}

func getAllRealEstates(context *gin.Context) {
	realEstates := models.GetAllRealEstate()
	context.JSON(http.StatusOK, realEstates)
}
