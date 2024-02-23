package main

import (
	"fmt"
	"os"

	"github.com/CrossStack-Q/Go_Backend_2024/tree/main/Projects/src/go-jwt-gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRouter(router)
	routes.UserRouter(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access Grated for Api 1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access Grated for Api 2"})
	})

	router.Run(":" + port)

	fmt.Println("Let's go with GO")

}
