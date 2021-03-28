package database

import (
	"goreact/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:root@/go_admin"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = database

	database.AutoMigrate(&models.User{}, &models.Role{})
}
