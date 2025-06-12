// db.go
package main

import (
	"gotaskify/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("/app/data/todos.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	err = DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("failed to auto-migrate: ", err)
	}
}
