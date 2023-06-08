package main

import "github.com/gin-gonic/gin"

func main() {
	//Gin实例
	r := gin.Default()
	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:8023")
}
