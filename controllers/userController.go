package controllers

import (
	"github.com/REM/initializers"
	"github.com/REM/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func SaveUser(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	result := initializers.DB.Create(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Błąd podczas dodawania nowego użytkownika"})
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message":    "Dodano nową nieruchomość",
		"Użytkownik": result,
	})
}

func Signup(c *gin.Context) {
	var body struct {
		Email     string
		Password  string
		FirstName string
		LastName  string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Błędne lub niekompletne dane",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Błąd podczas szyfrowania hasła",
		})
		return
	}

	user := models.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  string(hash),
		Enabled:   true,
		Locked:    false,
	}

	result := initializers.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"message":    "Pomyślnie utworzono użytkownika",
		"Użytkownik": result,
	})
}
