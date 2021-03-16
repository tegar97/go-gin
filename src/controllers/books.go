package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rest/src/config"
	"go-rest/src/models"
	"net/http"
)

func FindByBooks(c *gin.Context) {
	var books []models.Book
	config.DB.Find(&books)
	fmt.Println(books)

	c.JSON(http.StatusOK, gin.H{"data": books})

}
