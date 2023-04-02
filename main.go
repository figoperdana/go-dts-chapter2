package main

import (
	"log"

	"tugas8/config"
	"tugas8/models"
	"tugas8/repository"
	"tugas8/routers"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auto-migrate model Book
	if err := db.AutoMigrate(&models.Book{}).Error; err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}

	repo := repository.NewBookRepository(db)

	r := routers.SetupRouter(repo)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}