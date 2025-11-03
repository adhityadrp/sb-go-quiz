package main

import (
	"fmt"
	"log"
	"os"
	"sb-go-quiz/config"
	"sb-go-quiz/middlewares"
	"sb-go-quiz/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Tidak dapat memuat .env, menggunakan default environment")
	}

	// Init database & JWT
	config.InitDB()
	middlewares.InitJWT()

	// Run database migrations
	if err := config.RunMigrations(); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Jalankan server
	r := routes.SetupRouter()

	// Railway menggunakan PORT environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("APP_PORT") // fallback ke APP_PORT untuk development
		if port == "" {
			port = "8080" // default port
		}
	}

	log.Printf("Server starting on port %s", port)
	r.Run(fmt.Sprintf(":%s", port))
}
