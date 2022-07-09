package config

import (
	"app/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=ec2-34-239-241-121.compute-1.amazonaws.com user=gcbuxphbdrheuy password=9fbe7f4aad460c3b93ca1cb97fee21dfd4504c5059475940d06f39baa9cc1b51 dbname=dcvndkl01j4frj port=5432 sslmode=require TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed connect database !" + err.Error())
	}

	// defer db.Close()

	fmt.Println("* * * * * * * * * ")
	fmt.Println("Database Connected")
	fmt.Println("* * * * * * * * * ")

	// Migrate Schema
	db.AutoMigrate(&models.User{})

	DB = db
}
