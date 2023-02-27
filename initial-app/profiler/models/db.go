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

	if err = db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
	log.Println("Migrated User Table")

	if err = db.AutoMigrate(&Card{}); err != nil {
		panic(err)
	}
	log.Println("Migrated Card Table")

	if err = db.AutoMigrate(&CardType{}); err != nil {
		panic(err)
	}
	log.Println("Migrated CardType Table")

	DB = db
}
