package routes

import (
	"startupdigital/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/register", controller.Register)
	app.Post("/login", controller.Login)
	app.Get("/user", controller.User)
	app.Post("/logout", controller.Logout)
	app.Post("/test", controller.JawabTest)

}
