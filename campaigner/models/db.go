package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(dbConnString string) {
	db, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err = db.AutoMigrate(&Campaign{}); err != nil {
		panic(err)
	}
	log.Println("Migrated Campaign Table")

	DB = db
}
