package main

import (
	"log"

	"github.com/Alitadie/fengfengzhidaodoc/models"
	"github.com/Alitadie/fengfengzhidaodoc/router"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("doc.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug()
	db.AutoMigrate(&models.User{}, &models.Book{}, &models.Chapter{}, &models.Document{})

	r := router.SetupRouter(db)
	r.Run(":8000")
}
