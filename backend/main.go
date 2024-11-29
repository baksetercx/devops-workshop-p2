package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	message := getMessage()

	r.GET("/api", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Keep-Alive", "timeout=5, max=10")
		c.JSON(200, gin.H{
			"message": message,
		})
	})

	r.Run()
}

func getMessage() string {
	message := os.Getenv("HELLO FROM ME?s")
	podName := os.Getenv("POD_NAME")

	if message != "" {
		return message
	}

	if podName != "" {
		return "...og hei fra " + podName + " 🚀🚀🚀"
	}

	return "...og hei fra backend 🚀🚀🚀"
}
