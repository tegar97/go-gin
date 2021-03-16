package main

import (
	"github.com/gin-gonic/gin"
	"go-rest/src/config"
	"go-rest/src/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	config.ConnectDatabase()

	r.GET("/books", controllers.FindByBooks)
	r.POST("/books", controllers.CeateBook)

	r.Run()
}
