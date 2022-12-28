package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shh26b/go-code/fiber_mongo/db"
)

func createBook(c *fiber.Ctx) error {
	b := book{}
	err := c.BodyParser(&b)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	var status int
	status, err = db.CreateOne(c.Context(), bookColl, &b)
	if err != nil {
		return c.Status(status).SendString(err.Error())
	}
	return c.Status(status).JSON(&b)
}

func getBooks(c *fiber.Ctx) error {
	books := []book{}
	status, err := db.GetMany(c.Context(), bookColl, &books, msa{})
	if err != nil {
		return c.Status(status).SendString(err.Error())
	}
	return c.Status(status).JSON(&books)
}

func getBook(c *fiber.Ctx) error {
	id := c.Params("id")
	b := book{}
	status, err := db.GetOne(c.Context(), bookColl, &b, id)
	if err != nil {
		return c.Status(status).SendString(err.Error())
	}
	return c.Status(status).JSON(&b)
}

func deleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	status, err := db.DeleteOne(c.Context(), bookColl, id)
	if err != nil {
		return c.Status(status).SendString(err.Error())
	}
	return c.Status(status).JSON(msa{
		"msg": "Book with " + id + " successfully deleted",
	})
}

func updateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	b := book{}
	err := c.BodyParser(&b)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	status, err := db.UpdateOne(c.Context(), bookColl, &b, id)
	if err != nil {
		return c.Status(status).SendString(err.Error())
	}
	return c.Status(status).JSON(&b)
}
