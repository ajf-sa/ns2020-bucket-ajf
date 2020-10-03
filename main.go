package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views", ".html")
	// engine.Reload(true)    // Optional. Default: false
	// engine.Debug(true)     // Optional. Default: false
	engine.Layout("embed") // Optional. Default: "embed"
	app := fiber.New(fiber.Config{
		Views:     engine,
		BodyLimit: 1000 * 1024 * 1024,
		// Prefork:   true,
	})
	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{"Title": "Public"}, "layout")
	})
	app.Post("/", func(ctx *fiber.Ctx) error {
		file, _ := ctx.FormFile("fileToUpload")
		ctx.SaveFile(file, fmt.Sprintf("./public/%s", file.Filename))
		return ctx.Redirect("/", 302)
	})
	app.Listen(":3000")
}
