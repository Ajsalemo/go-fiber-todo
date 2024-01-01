package controllers

import (
	"encoding/json"
	config "go-fiber-todo-backend/config"
	models "go-fiber-todo-backend/models"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func CreateTodo(cxt *fiber.Ctx) error {
	var tasks models.Task

	db, err := config.ConnectDB()
	// This helper function implements exponential retry backoffs for connection failure attempts to the database
	config.ConnectionRetry(err)

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
