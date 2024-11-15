package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thutaminthway/go-fiber-gorm/internal/controllers"
	"github.com/thutaminthway/go-fiber-gorm/internal/middleware"
)

func Setup(app *fiber.App) {
	api := app.Group("api")

	SetupAuthRoutes(api)
	SetupUserRoutes(api)
	SetupBookRoutes(api)
}

func SetupUserRoutes(api fiber.Router) {
	user := api.Group("users")
	user.Get("/", controllers.GetUsers)
	user.Get("/:id", controllers.ShowUser)
	user.Post("/", controllers.CreateUser)
	user.Put("/:id", controllers.UpdateUser)
	user.Delete("/:id", controllers.DeleteUser)
}

func SetupBookRoutes(api fiber.Router) {
	book := api.Group("books")
	book.Get("/", controllers.GetBooks)
	book.Get("/:id", controllers.ShowBook)
	book.Post("/", controllers.CreateBook)
	book.Put("/:id", controllers.UpdateBook)
	book.Delete("/:id", controllers.DeleteBook)
}

func SetupAuthRoutes(api fiber.Router) {
	auth := api.Group("auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
	auth.Get("/me", middleware.Protected(), controllers.Me)
}
