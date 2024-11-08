package migration

import (
	"github.com/thutaminthway/go-fiber-gorm/internal/database"
	"github.com/thutaminthway/go-fiber-gorm/internal/model"
)

func Migrate() error {
	return database.DB.AutoMigrate(&model.User{}, &model.Book{})
}
