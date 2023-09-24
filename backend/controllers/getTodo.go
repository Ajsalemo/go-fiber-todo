package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetTodo(cxt *fiber.Ctx) error {
	return cxt.JSON(fiber.Map{"msg": "gettodo"})
}
