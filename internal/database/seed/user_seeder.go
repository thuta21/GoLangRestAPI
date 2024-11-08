package seed

import (
	"log"

	"github.com/thutaminthway/go-fiber-gorm/internal/database"
	"github.com/thutaminthway/go-fiber-gorm/internal/model"
)

func SeedUsers() error {
	users := []model.User{
		{Name: "Alice", Email: "alice@example.com"},
		{Name: "Bob", Email: "bob@example.com"},
	}

	for _, user := range users {
		if err := database.DB.Create(&user).Error; err != nil {
			log.Printf("Failed to seed user: %v\n", err)
			return err
		}
	}

	log.Println("User seeding completed successfully!")
	return nil
}
