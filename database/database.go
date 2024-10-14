package database

import (
	"log"

	"github.com/herumitra/fiberapi.git/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:@tcp(localhost:3306)/goapi?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database")
	}

	DB = connection

	DB.AutoMigrate(&models.User{}, &models.LogFailure{}, &models.Branch{}, &models.TokenBlacklist{}, &models.Unit{})
}
