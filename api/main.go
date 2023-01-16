package main 

import (
	"github.com/gofiber/fiber/v2"
)

func getIndex(c *fiber.Ctx) error {
	return c.SendString("Hello, Chinese Flashcard API.");
}

func main() {
	app := fiber.New();
	app.Get("/", getIndex);
	app.Listen(":8002"); 
}