package controllers

import (
	"fmt"
	"net/http"
	"os"

	// "os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	// "github.com/golang-jwt/jwt"
	initializers "github.com/nidhey27/go-jwt/initilizers"
	"github.com/nidhey27/go-jwt/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "",
			"data":    make([]string, 0),
			"error":   "Failed to read REQUEST Body",
		})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "",
			"error":   "Failed to Hash Password",
			"data":    make([]string, 0),
		})
		return
	}

	// Create the user
	user := models.User{Email: body.Email, Password: string(passwordHash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "",
			"error":   result.Error,
			"data":    make([]string, 0),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "User Created!",
		"error":   "",
		"data":    user,
	})

}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "",
			"data":    make([]string, 0),
			"error":   "Failed to read REQUEST Body",
		})
		return
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "",
			"data":    make([]string, 0),
			"error":   "Invalid Email ID",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "",
			"error":   "Invalid Password",
			"data":    make([]string, 0),
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	fmt.Println(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "TOKEN GENERATION",
			"error":   err.Error(),
			"data":    make([]string, 0),
		})
		return
	}

	// NOT WORKING
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	// NOT WORKING

	// c.Header("Authorization", tokenString)
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Login Successful!",
		"error":   "",
		"data":    make([]string, 0),
	})

}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Authorized",
		"error":   "",
		"data":    make([]string, 0),
	})
}
