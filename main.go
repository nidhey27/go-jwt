package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nidhey27/go-jwt/controllers"
	initializers "github.com/nidhey27/go-jwt/initilizers"
	"github.com/nidhey27/go-jwt/middlewares"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "GO JWT Authentication - POSTGRES API",
			"error":   "",
			"author":  "Nidhey Indurkar",
			"data":    make([]string, 0),
		})
	})

	r.POST("/sign-up", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middlewares.ValidateJWT, controllers.Validate)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
