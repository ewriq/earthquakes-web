package main

import (
	"github.com/gofiber/fiber/v2"
	"deprem/Router"
)



func main() {
	app := fiber.New()
   

	Router.Initalize(app)
	

	app.Listen(":3000")
}
