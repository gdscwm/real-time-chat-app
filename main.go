package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html/v2"
	"github.com/jdkuffa/handlers"
)

func main() {

    // Create views engine
    viewsEngine := html.New("./views", ".html")

    // Start new fiber instance
    app := fiber.New(fiber.Config{
        Views: viewsEngine,
    })

    // Static route and directory
    app.Static("/static/", "./static")

    // Create a "ping" handler to test the server
    app.Get("/ping", func(ctx *fiber.Ctx) error{
        return ctx.SendString("Welcome to fiber")
    })

	// create new App Handler
    appHandler := NewAppHandler()

    // Add appHandler routes
    app.Get("/", appHandler.HandleGetIndex)

    // Start the http server
    app.Listen(":3000")

}