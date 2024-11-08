package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/thutaminthway/go-fiber-gorm/internal/database"
	"github.com/thutaminthway/go-fiber-gorm/internal/model"
	"github.com/thutaminthway/go-fiber-gorm/internal/utils"
)

func GetBooks(c *fiber.Ctx) error {
	var books []model.Book
	if result := database.DB.Order("id DESC").Find(&books); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch books"})
	}

	return c.JSON(books)
}

func ShowBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book model.Book

	if result := database.DB.First(&book, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	return c.JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	book := new(model.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := utils.Validate.Struct(book); err != nil {
		errors := err.(validator.ValidationErrors)
		var errorMessages []string
		for _, e := range errors {
			errorMessages = append(errorMessages, e.Field()+" failed validation: "+e.ActualTag())
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errorMessages})
	}

	if result := database.DB.Create(&book); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create book"})
	}

	return c.Status(fiber.StatusCreated).JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book model.Book

	if result := database.DB.First(&book, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := utils.Validate.Struct(&book); err != nil {
		errors := err.(validator.ValidationErrors)
		var errorMessages []string
		for _, e := range errors {
			errorMessages = append(errorMessages, e.Field()+" failed validation: "+e.ActualTag())
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errorMessages})
	}

	if result := database.DB.Save(&book); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update book"})
	}

	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book model.Book

	if result := database.DB.First(&book, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	if result := database.DB.Delete(&book); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete book"})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"message": "Book deleted successfully"})
}
