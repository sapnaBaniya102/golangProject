package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main()  {

	engine := html.New("./public",".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public/dist")

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{})
	})

	app.Listen(":3000")
}
