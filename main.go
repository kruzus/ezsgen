package main

import (
	"log"

	handler "ezsgen/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	// Print single password
	app.Get("/", handler.Root)

	// Generate more passwords
	app.Get("/more", handler.Multiple)

	// DEV MODE ONLY
	// app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	log.Fatal(app.Listen(":3000"))

}
