package main

import (
	"github.com/REM/controllers"
	"github.com/REM/initializers"
	"github.com/REM/middleware"
	"github.com/REM/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	server := gin.Default()

	//CORS config
	config := cors.DefaultConfig()
	config.AllowOrigins = append(config.AllowOrigins, "http://localhost:5173")
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	server.Use(cors.New(config))

	server.GET("/real-estates", getAllRealEstates)
	server.POST("/real-estates", controllers.SaveRealEstate)
	server.POST("/user/signup", controllers.Signup)
	server.POST("/login", controllers.Login)
	server.GET("/validate", middleware.RequireAuth, controllers.Validate)

	err := server.Run()
	if err != nil {
		return
	}
}

func getAllRealEstates(context *gin.Context) {
	realEstates := models.GetAllRealEstate()
	context.JSON(http.StatusOK, realEstates)
}
