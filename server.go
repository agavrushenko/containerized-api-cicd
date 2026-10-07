package main

import (
    "log"
    "time"
    "encoding/json"
    "github.com/gofiber/fiber/v3"
)

func main() {
    app := fiber.New(fiber.Config{
		JSONEncoder: func(v interface{}) ([]byte, error) {
			return json.MarshalIndent(v, "", "  ")
		},
	})


    app.Get("/", func(c fiber.Ctx) error {
        return c.JSON(fiber.Map{
	    "message": "My name is Alexey Gavrushenko", 
	    "timestamp": time.Now().UnixMilli(),
        })
    })

    log.Fatal(app.Listen(":8080"))
}
