package main

import (
	"auth-server-go/controller"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	app := fiber.New()

	app.Post("/auth-service/api/v1/auth", controller.Auth)

	log.Fatal(app.Listen(":9006"))
}
