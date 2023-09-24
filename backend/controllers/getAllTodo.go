package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllTodo(cxt *fiber.Ctx) error {
	return cxt.JSON(fiber.Map{"msg": "getalltodo"})
}
