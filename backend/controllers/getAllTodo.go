package controllers

import (
	"encoding/json"
	config "go-fiber-todo-backend/config"
	models "go-fiber-todo-backend/models"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func GetAllTodo(cxt *fiber.Ctx) error {
	var tasks []models.Task

	db, err := config.ConnectDB()
	// This helper function implements exponential retry backoffs for connection failure attempts to the database
	config.ConnectionRetry(err)

	db.Find(&tasks)
	// For now, marshal the return tasks struct into JSON and parse into a sring
	out, err := json.Marshal(&tasks)
	if err != nil {
		zap.L().Error(err.Error())
	}

	zap.L().Info(string(out))
	// No tasks found is returned as an empty array
	// Let the client-side handle displaying of no tasks, if so
	return cxt.JSON(fiber.Map{"data": &tasks})
}
