package service

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lld-tdd/models"
	"log"
)

func ConnectDB() (*gorm.DB, error) {
	// Database connection configuration
	dsn := "root:@tcp(localhost:3306)/lldtdd?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	log.Println("Database migration completed..")
}
