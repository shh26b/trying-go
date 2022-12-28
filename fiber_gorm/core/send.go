package core

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Err struct {
	C      *fiber.Ctx
	DB     *gorm.DB
	Status int
	Msg    string
}

func (e Err) Send() error {
	log.Println(e.Msg, ": -> ", e.DB.Error)
	e.C.Status(e.Status)
	return e.C.JSON(map[string]any{"msg": e.Msg})
}
