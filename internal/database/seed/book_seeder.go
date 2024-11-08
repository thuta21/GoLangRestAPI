package seed

import (
	"log"

	"github.com/thutaminthway/go-fiber-gorm/internal/database"
	"github.com/thutaminthway/go-fiber-gorm/internal/model"
)

func SeedBooks() error {
	books := []model.Book{
		{Title: "Go Programming", Author: "John Doe", Description: "An introduction to Go programming."},
		{Title: "Advanced Fiber", Author: "Jane Smith", Description: "Deep dive into Fiber framework."},
	}

	for _, book := range books {
		if err := database.DB.Create(&book).Error; err != nil {
			log.Printf("Failed to seed book: %v\n", err)
			return err
		}
	}

	log.Println("Book seeding completed successfully!")
	return nil
}
