package main

import (
	"context"
	"embed"
	"log"

	"github.com/a-h/templ"
	"github.com/dfalgout/gofiber/api"
	"github.com/dfalgout/gofiber/ent"
	"github.com/dfalgout/gofiber/pages"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed assets/*
var assets embed.FS

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	app := fiber.New(fiber.Config{
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		DisableStartupMessage: true,
	})

	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(recover.New())

	// Favicon middleware
	app.Use(favicon.New(favicon.Config{
		File: "./favicon.ico",
		URL:  "/favicon.ico",
	}))

	app.Static("/assets", "./assets", fiber.Static{
		Compress: true,
	})

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins: "https://test.com",
	}))

	app.Use(func(c *fiber.Ctx) error {
		if c.Route().Method == fiber.MethodGet && c.Route().Path == "/" {
			return c.Next()
		}
		// user := c.Cookies("token")
		// if user == "" {
		// 	return c.Redirect("/")
		// }
		// ctx := context.Background()
		// userContext := auth.Add(ctx, user)
		// c.SetUserContext(userContext)
		return c.Next()
	})

	// view routes
	pages.Register(app)

	// api routes
	api.Register(app, client)

	// register global css middleware for templ css classes
	// Needs to be the last route since next doesn't adapt from fiber
	bodyStyle := pages.BodyStyle()
	m := templ.NewCSSMiddleware(nil, bodyStyle.(templ.ComponentCSSClass))
	app.Use(adaptor.HTTPHandler(m))

	log.Printf("server started at http://127.0.0.1:3000")
	log.Fatal(app.Listen(":3000"))
}
