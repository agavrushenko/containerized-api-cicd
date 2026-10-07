package main

import (
    "log"
    "time"
    "github.com/gofiber/fiber/v3"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c fiber.Ctx) error {
        return c.JSON(fiber.Map{
	    "message": "My name is Alexey Gavrushenko", 
	    "timestamp": time.Now().UnixMilli(),
	    "test-deloy": "Success",
        })
    })

    log.Fatal(app.Listen(":8080"))
}
