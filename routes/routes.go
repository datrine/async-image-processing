package routes

import (
	"github.com/datrine/async-image-processing/handlers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/image-process", handlers.UploadImage)
}
