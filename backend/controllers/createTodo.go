package controllers

import (
	"encoding/json"
	models "go-fiber-todo-backend/models"
	config "go-fiber-todo-backend/config"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func CreateTodo(cxt *fiber.Ctx) error {
	var tasks models.Task

	db, err := config.ConnectDB()
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	body := cxt.Body()
	// Return an error if the request body is empty
	if body == nil {
		zap.L().Error("Request body is nil")
		return cxt.Status(400).JSON(fiber.Map{"msg": "Request body is empty"})
	}
	// Unmarshal the JSON into a Task object
	// Return an error if the JSON is invalid
	err2 := json.Unmarshal(body, &tasks)
	if err2 != nil {
		zap.L().Error(err.Error())
	}
	// Create a new task
	db.Create(&tasks)
	zap.L().Info("Created a new task: " + string(body))
	return cxt.JSON(fiber.Map{"msg": "createtodo"})
}
