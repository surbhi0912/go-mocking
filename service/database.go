package service

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lld-tdd/models"
	"log"
)

var DBInstance *gorm.DB
var err error

func Connect() {
	dsn := "root:@tcp(localhost:3306)/lldtdd?charset=utf8mb4&parseTime=True&loc=Local"
	DBInstance, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to DB")
}

func Migrate() {
	// creates table if it doesn't exist
	DBInstance.AutoMigrate(&models.User{})
	log.Println("Database migration completed..")
}
