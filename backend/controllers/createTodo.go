package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func CreateTodo(cxt *fiber.Ctx) error {
	return cxt.JSON(fiber.Map{"msg": "createtodo"})
}
