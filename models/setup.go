package models

// import dependency & library
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Declare variabel DB in models package
var DB *gorm.DB

// Create function ConnectDB
func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:Aliya01072023_@tcp(localhost:3306)/go_rest_fiber"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Book{})
	DB = db
}
