package main

import (
	"fmt"
	"log"

	"currency_converter/config"
	"currency_converter/database"
	"currency_converter/handlers"
)

func main() {
	config.LoadConfig()

	app := fiber.New()

	if err := database.Connect(); err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	app.Post("/convert", handlers.Convert)

	port := config.Cfg.Server.Port
	log.Printf("Server is running on port %d", port)
	if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
