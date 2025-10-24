package main

import (
	"log"
	"ranggaAdiPratama/go-with-claude/internal/config"
	"ranggaAdiPratama/go-with-claude/internal/database"
	"ranggaAdiPratama/go-with-claude/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	cfg := config.Load()

	db, err := database.NewConnection(cfg.DatabaseURL)

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	defer db.Close()

	srv := server.New(db, cfg)

	log.Printf("Server starting on port %s", cfg.Port)

	if err := srv.Start(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
