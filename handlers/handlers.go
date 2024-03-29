package handlers

import (
	"fmt"
	"io"

	"github.com/datrine/async-image-processing/tasks"
	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) error {

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Upload failed"})
	}

	fileData, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to open the file"})
	}
	defer fileData.Close()

	data, err := io.ReadAll(fileData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to read the file"})
	}

	resizeTasks, err := tasks.NewImageResizeTasks(data, file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create image resize tasks"})
	}
	client := tasks.GetClient()
	for _, task := range resizeTasks {
		if _, err := client.Enqueue(task); err != nil {
			fmt.Printf("Error enqueuing task: %v\n", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not enqueue image resize task"})
		}
	}
	return c.JSON(fiber.Map{"message": "Image uploaded and resizing tasks started"})
}
