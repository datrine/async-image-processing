package main

import (
	"log"

	"github.com/datrine/async-image-processing/routes"
	"github.com/datrine/async-image-processing/tasks"
	"github.com/gofiber/fiber/v2"
)

const redisAddress = "127.0.0.1:6379"

func main() {
	app := fiber.New()
	routes.Setup(app)
	tasks.Init(redisAddress)
	defer tasks.Close()
	log.Fatal(app.Listen(":3000"))
}
