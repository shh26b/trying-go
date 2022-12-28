package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/shh26b/go-code/fiber_gorm/book"
	"github.com/shh26b/go-code/fiber_gorm/core"
)

func setV1(c *fiber.Ctx) error {
	c.Set("Version", "v1")
	err := c.Next()
	return err
}

func setUpRouters(app *fiber.App) {
	apiV1 := app.Group("/api/v1", setV1)
	book.Router(apiV1.Group("/books"))
}

func initDB() *sql.DB {
	var err error
	core.DB, err = gorm.Open(sqlite.Open("tmp.db"), &gorm.Config{})
	core.Fatal(err)
	log.Println("Database connection successfully connected")

	core.DB.AutoMigrate(&book.Book{})
	log.Println("Database book table migrated successfully")

	sqlDB, err := core.DB.DB()
	core.Fatal(err)
	return sqlDB
}

func main() {
	app := fiber.New()

	sqlDB := initDB()
	defer sqlDB.Close()

	app.Use(logger.New())
	setUpRouters(app)

	app.Listen(":3000")
}
