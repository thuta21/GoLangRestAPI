package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/thutaminthway/go-fiber-gorm/internal/config"
	"github.com/thutaminthway/go-fiber-gorm/internal/database"
	"github.com/thutaminthway/go-fiber-gorm/internal/database/migration"
	"github.com/thutaminthway/go-fiber-gorm/internal/database/seed"
	"github.com/thutaminthway/go-fiber-gorm/internal/routes"
)

func main() {
	migrate := flag.Bool("migrate", false, "Apply database migrations")
	seedData := flag.Bool("seed", false, "Seed database with dummy data")
	flag.Parse()

	cfg := config.LoadConfig()

	_, err := database.Init(cfg)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	if *migrate {
		fmt.Println("Applying migrations...")
		if err := migration.Migrate(); err != nil {
			log.Fatalf("Failed to apply migrations: %v", err)
		}
		fmt.Println("Migrations applied successfully!")
		return
	}

	if *seedData {
		fmt.Println("Seeding database with dummy data...")
		if err := seed.Seed(); err != nil {
			log.Fatalf("Failed to seed database: %v", err)
		}
		fmt.Println("Seeding completed successfully!")
		return
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://gofiber.io, https://gofiber.net",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.Setup(app)

	log.Println("Server running on http://localhost:3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
