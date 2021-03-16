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

func CeateBook(c *gin.Context) {
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	config.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBook(c *gin.Context) { // Get model if exist
	var book models.Book

	if err := config.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func updateBook(c *gin.Context) {
	var book models.Book
	if err := config.DB.Where("id= ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.UpdateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})

}
