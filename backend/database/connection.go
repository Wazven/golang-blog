package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/wazven/backendblog/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB

func Connection() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Load .env file")
	}

	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn),  &gorm.Config{})
	
	if err != nil {
		panic("Can't Connect to the Database")
	} else {
		log.Println("Connect Successfully")
	}
	DB = database
	database.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)
}