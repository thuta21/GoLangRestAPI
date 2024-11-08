package seed

import "log"

func Seed() error {
	if err := SeedUsers(); err != nil {
		log.Println("Error seeding users:", err)
		return err
	}

	if err := SeedBooks(); err != nil {
		log.Println("Error seeding books:", err)
		return err
	}

	log.Println("All seeding completed successfully!")
	return nil
}
