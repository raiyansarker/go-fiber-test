package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

func main() {
	app := fiber.New()
	app.Use(func (c fiber.Ctx) error {
		c.Set("Server", "RSA")
		return c.Next()
	})
	app.Use(requestid.New())
	app.Use(logger.New())

	app.Get("/", func(c fiber.Ctx) error {
		slug := c.Params("slug")
		if (slug == "") {
			slug = "/"
		}

		now := time.Now().String()
		hostname, err := os.Hostname()
		if (err != nil) {
			log.Fatal("Hostname not found")
		}

		data, err := json.Marshal(struct {
			Hostname string `json:"hostname"`
			Time string `json:"time"`
		}{
			Hostname: hostname,
			Time: now,
		})
		if (err != nil) {
			log.Fatal("JSON Marshal failed")
		}

		output := fmt.Sprintf("<h1>Slug %v</h1><br /><pre>%v</pre>", slug, string(data))

		return c.AutoFormat(output)
	})

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}