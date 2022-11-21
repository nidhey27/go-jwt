package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	initializers "github.com/nidhey27/go-jwt/initilizers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
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

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
