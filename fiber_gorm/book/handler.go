package book

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/shh26b/go-code/fiber_gorm/core"
)

func getBooks(c *fiber.Ctx) error {
	books := []Book{}

	r := core.DB.Find(&books)
	if r.Error != nil {
		e := core.Err{C: c, DB: r, Status: 500, Msg: "Can't get all books in db"}
		return e.Send()
	}
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	id := c.Params("id")

	book := Book{}

	r := core.DB.First(&book, id)
	if book.Title == "" {
		m := "Can't get one book for id " + string(id) + " in db"
		e := core.Err{C: c, DB: r, Status: 404, Msg: m}
		return e.Send()
	}
	if r.Error != nil {
		m := "Can't get one book for id " + string(id) + " in db"
		e := core.Err{C: c, DB: r, Status: 500, Msg: m}
		return e.Send()
	}

	return c.JSON(book)
}

func createBook(c *fiber.Ctx) error {
	book := Book{}
	err := c.BodyParser(&book)
	if err != nil {
		e := core.Err{C: c, DB: core.DB, Status: 400, Msg: "Wrong data"}
		return e.Send()
	}

	r := core.DB.Create(&book)
	if r.Error != nil {
		e := core.Err{C: c, DB: r, Status: 500, Msg: "Can't create new book in db"}
		return e.Send()
	}

	c.Status(http.StatusCreated)
	return c.JSON(book)
}

func updateBook(c *fiber.Ctx) error {
	id := c.Params("id")

	book := Book{}

	r := core.DB.First(&book, id)
	if book.Title == "" {
		m := "Can't get one book for id " + string(id) + " in db"
		e := core.Err{C: c, DB: r, Status: 404, Msg: m}
		return e.Send()
	}
	if r.Error != nil {
		m := "Can't get one book for id " + string(id) + " in db"
		e := core.Err{C: c, DB: r, Status: 500, Msg: m}
		return e.Send()
	}

	err := c.BodyParser(&book)
	if err != nil {
		e := core.Err{C: c, DB: core.DB, Status: 400, Msg: "Wrong data"}
		return e.Send()
	}

	r = core.DB.Save(&book)
	if r.Error != nil {
		m := "Can't save in update of book for id " + string(id) + " in db"
		e := core.Err{C: c, DB: r, Status: 500, Msg: m}
		return e.Send()
	}

	return c.JSON(book)
}

func deleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := core.DB

	book := Book{}
	r := core.DB.First(&book, id)
	if book.Title == "" {
		m := "Can't get one book for id " + string(id) + " in db"
		e := core.Err{C: c, DB: r, Status: 404, Msg: m}
		return e.Send()
	}

	r = db.Delete(&book)
	if r.Error != nil {
		e := core.Err{C: c, DB: r, Status: 500, Msg: "Can't delete book in db"}
		return e.Send()
	}

	return c.JSON(map[string]any{
		"msg": "Book with " + id + " successfully deleted",
	})
}
