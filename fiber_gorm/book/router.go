package book

import (
	"github.com/gofiber/fiber/v2"
)

func m1(c *fiber.Ctx) error {
	return c.Next()
}

func m2(c *fiber.Ctx) error {
	return c.Next()
}

func Router(r fiber.Router) {
	r.Get("/:id", m1, m2, getBook)
	r.Patch("/:id", updateBook)
	r.Delete("/:id", deleteBook)
	r.Get("/", getBooks)
	r.Post("/", createBook)
}
