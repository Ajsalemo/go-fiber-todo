package controllers

import (
	"encoding/json"
	models "go-fiber-todo-backend/models"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func CreateTodo(cxt *fiber.Ctx) error {
	var tasks models.Task

	body := cxt.Body()
	zap.L().Info(string(body))
	err := json.Unmarshal(body, &tasks)
	if err != nil {
		zap.L().Error(err.Error())
	}

	return cxt.JSON(fiber.Map{"msg": "createtodo"})
}
