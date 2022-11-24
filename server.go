package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// It creates a new router.
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "testing OK!!",
		})
	})

	r.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "testing OK!! \n POST",
		})
	})

	// listen and serve on
	r.Run()
}
