package database

import (
	"github.com/jintonick/SigmaBank/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("sql11661721:jXpMVhlBuk@tcp(sql11.freemysqlhosting.net)/sql11661721"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	DB = connection

	err = connection.AutoMigrate(&models.User{}, &models.Point{}, &models.Task{})
	if err != nil {
		panic("could not start server")
	}
}
