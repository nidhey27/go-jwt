package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	initializers "github.com/nidhey27/go-jwt/initilizers"
	"github.com/nidhey27/go-jwt/models"
)

func ValidateJWT(c *gin.Context) {
	fmt.Println("Working")
	tokenString, err := c.Cookie("Authorization")
	fmt.Println(tokenString)
	if tokenString == "" || err != nil {
		fmt.Println("1")
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  true,
			"message": "UnAuthorized",
			"error":   "",
			"data":    make([]string, 0),
		})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println(ok)
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		fmt.Println(os.Getenv("JWT_SECRET"))
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		fmt.Println("2")
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  true,
			"message": "UnAuthorized",
			"error":   "",
			"data":    make([]string, 0),
		})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			fmt.Println("3")
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  true,
				"message": "UnAuthorized",
				"error":   "",
				"data":    make([]string, 0),
			})
			return
		}

		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			fmt.Println("4")
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  true,
				"message": "UnAuthorized",
				"error":   "",
				"data":    make([]string, 0),
			})
			return
		}

		c.Set("user", user)
		c.Next()
	} else {
		fmt.Println("5")

		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  true,
			"message": "UnAuthorized",
			"error":   "",
			"data":    make([]string, 0),
		})
		return
	}

}
