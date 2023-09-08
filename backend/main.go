package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/wazven/backendblog/database"
	"github.com/wazven/backendblog/routes"
)

func main() {
	//Connect Database
	database.Connection()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env File")
	}
	
	app:= fiber.New()
	routes.Setup(app)


	port := os.Getenv("PORT")
	app.Listen(":"+port)
}