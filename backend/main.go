package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/sethvargo/go-password/password"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	app.Get("/genpass", func(c *fiber.Ctx) error {
		gen := makePassword(64, 10, 10, false, false)
		return c.JSON(fiber.Map{
			"msg":                 "Welcome",
			"genereated_password": gen,
		})

	})

	app.Get("/genpasses", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"msg":                  "More Options",
			"genereated_passwords": generate_multiple(),
		})
	})

	log.Fatal(app.Listen(":9090"))

}

/*
# makePassword

Takes: length, numDigits, numSymbols int, noUpper, allowRepeat bool
*/
func makePassword(length, numDigits, numSymbols int, noUpper, allowRepeat bool) string {
	res, err := password.Generate(length, numDigits, numSymbols, noUpper, allowRepeat)
	if err != nil {
		log.Fatal(err)
	}
	return res

}

// @TODO: Work on this function and make it custum for the api
func generate_multiple() []string {
	var data []string
	for i := 1; i < 10; i++ {
		data = append(data, makePassword(64, 10, 10, false, false))
	}
	return data
}
