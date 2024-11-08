package model

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string `json:"title" validate:"required"`
	Author      string `json:"author" validate:"required"`
	Description string `json:"description" validate:"required"`
}
