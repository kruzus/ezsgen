package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/sethvargo/go-password/password"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		gen := makePassword(64, 10, 10, false, false)
		return c.JSON(fiber.Map{
			"msg":                 "Welcome",
			"genereated_password": gen,
		})

	})

	app.Get("/more", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"msg":                  "More Options",
			"genereated_passwords": generate_multiple(),
		})
	})

	log.Fatal(app.Listen(":3000"))

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
