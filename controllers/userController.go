package controllers

import (
	"github.com/REM/initializers"
	"github.com/REM/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
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

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Użytkownik z podanym e-mailem już istnieje!",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Pomyślnie utworzono użytkownika",
		"Użytkownik": result,
	})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Błąd podczas procesowania zapytania logowania",
		})
		return
	}

	var user models.User

	initializers.DB.First(&user, "email=?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Błędny email lub hasło",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Błędny email lub hasło",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	secret := os.Getenv("SECRET")

	tokenString, tokenErr := token.SignedString([]byte(secret))

	if tokenErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Błąd podczas generowania tokenu jwt",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
