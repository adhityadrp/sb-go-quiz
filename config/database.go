package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Tidak bisa memuat file .env, menggunakan environment default")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode,
	)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("❌ Gagal membuka koneksi ke database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("❌ Tidak bisa konek ke database:", err)
	}

	log.Println("✅ Database connected")
}

// RunMigrations executes all SQL migration files in the migrations directory
func RunMigrations() error {
	log.Println("🔄 Running database migrations...")

	files, err := filepath.Glob("migrations/*.sql")
	if err != nil {
		return fmt.Errorf("failed to find migration files: %v", err)
	}

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %v", file, err)
		}

		// Extract only the Up section between the migration markers
		text := string(content)
		upMarker := "-- +migrate Up"
		downMarker := "-- +migrate Down"

		upIdx := strings.Index(text, upMarker)
		if upIdx == -1 {
			log.Printf("skipping file (no Up marker): %s", file)
			continue
		}
		// start after the Up marker line
		upStart := upIdx + len(upMarker)

		downIdx := strings.Index(text, downMarker)
		var upSQL string
		if downIdx == -1 {
			// if no Down marker, take rest of file
			upSQL = text[upStart:]
		} else {
			upSQL = text[upStart:downIdx]
		}

		upSQL = strings.TrimSpace(upSQL)
		if upSQL == "" {
			log.Printf("no Up SQL to execute in %s", file)
			continue
		}

		log.Printf("Executing migration (Up): %s", file)

		_, err = DB.Exec(upSQL)
		if err != nil {
			return fmt.Errorf("failed to execute migration %s: %v", file, err)
		}
	}

	log.Println("✅ Migrations completed successfully")
	return nil
}
