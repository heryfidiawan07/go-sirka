package config

import (
	"app/models"
	"fmt"
	"os"

	"github.com/subosito/gotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	gotenv.Load()

	// dsn := "host=ec2-34-239-241-121.compute-1.amazonaws.com user=gcbuxphbdrheuy password=9fbe7f4aad460c3b93ca1cb97fee21dfd4504c5059475940d06f39baa9cc1b51 dbname=dcvndkl01j4frj port=5432 sslmode=require TimeZone=Asia/Jakarta"
	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=require TimeZone=Asia/Jakarta"
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
