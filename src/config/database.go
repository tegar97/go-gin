package config

import (
	"go-rest/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root@tcp(127.0.0.1:3306)/go_book?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}
	database.AutoMigrate(&models.Book{})

	DB = database

}
