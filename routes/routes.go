package routes

import (
	"startupdigital/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/register", controller.Register)
	app.Post("/login", controller.Login)

}
