package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/shh26b/go-code/fiber_mongo/book"
	"github.com/shh26b/go-code/fiber_mongo/db"
	"github.com/shh26b/go-code/fiber_mongo/util"
)

func setV1(c *fiber.Ctx) error {
	c.Set("Version", "v1")
	err := c.Next()
	return err
}

func setUpRouters(app *fiber.App) {
	apiV1 := app.Group("/api/v1", setV1)
	book.Router(apiV1)
}

func main() {

	app := fiber.New()

	var dBName = "fiber-mongo"
	var mongoURI = "mongodb://localhost:27017"

	err := db.Connect(dBName, mongoURI)
	util.Fatal(err)
	defer db.Disconnect()

	app.Use(logger.New())
	setUpRouters(app)

	app.Listen(":3000")
}
