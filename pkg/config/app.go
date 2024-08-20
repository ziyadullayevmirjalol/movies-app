package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	// Update the DSN format according to PostgreSQL
	dsn := "host=localhost user=postgres password=sadperson23 dbname=bookstore_go port=5432 sslmode=disable TimeZone=Asia/Tashkent"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
