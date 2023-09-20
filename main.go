package main

import (
	"context"
	"database/sql"
	"embed"
	"log"
	"path"

	"github.com/a-h/templ"
	"github.com/dfalgout/gofiber/api"
	"github.com/dfalgout/gofiber/dal"
	"github.com/dfalgout/gofiber/pages"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed assets/*
var assets embed.FS

//go:embed schemas
var schemas embed.FS

func main() {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	// auto migrations
	dirs, err := schemas.ReadDir("schemas")
	if err != nil {
		log.Fatal(err)
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			continue
		}
		file, err := schemas.ReadFile(path.Join("schemas", dir.Name()))
		if err != nil {
			log.Fatal(err)
		}
		if _, err := db.ExecContext(ctx, string(file)); err != nil {
			log.Fatal(err)
		}
	}

	queries := dal.New(db)

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

	// view routes
	pages.Register(app)

	// api routes
	api.Register(app, queries)

	// register global css middleware for templ css classes
	// Needs to be the last route since next doesn't adapt from fiber
	bodyStyle := pages.BodyStyle()
	m := templ.NewCSSMiddleware(nil, bodyStyle.(templ.ComponentCSSClass))
	app.Use(adaptor.HTTPHandler(m))

	log.Printf("server started at http://127.0.0.1:3000")
	log.Fatal(app.Listen(":3000"))
}
