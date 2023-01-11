package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:6995@tcp(localhost:3306)/dulrestful"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})
	database.AutoMigrate(&Transaction{})
	database.AutoMigrate(&User{})

	DB = database
}
