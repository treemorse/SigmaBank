package database

import (
	"github.com/jintonick/SigmaBank/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:Amaj7A6Am6G6@/auth"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	DB = connection

	err = connection.AutoMigrate(&models.User{}, &models.Point{}, &models.Task{})
	if err != nil {
		panic("could not start server")
	}
}
