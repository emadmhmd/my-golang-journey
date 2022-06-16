package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB = nil
var err error

func Connection() {
	dsn := "host=localhost user=postgres password=0000 dbname=Micro port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Fail to connect to postgres db")
	}
}

func Migrate() {
	db.AutoMigrate(&Token{})
}
