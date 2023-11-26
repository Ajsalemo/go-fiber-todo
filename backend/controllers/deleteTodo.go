package controllers

import (
	config "go-fiber-todo-backend/config"
	models "go-fiber-todo-backend/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func DeleteTodo(cxt *fiber.Ctx) error {
	id := cxt.Params("id")
	var tasks []models.Task

	db, err := config.ConnectDB()
	if err != nil {
		zap.L().Fatal(err.Error())
	}
	// If the request param is empty, return it to the client for potential error handling
	if id == "" {
		zap.L().Warn("Parameter must not be empty")
		return cxt.JSON(fiber.Map{"err": "Parameter must not be empty"})
	} else {
		// Parse the request param into an Int type
		parsedId, err := strconv.ParseInt(id, 10, 64)
		// If there is an issue with Int conversion with the request param, return it to the client for potential error handling
		if err != nil {
			zap.L().Error(err.Error())
			return cxt.JSON(fiber.Map{"err": err})
		}

		db.Find(&tasks, parsedId)

		db.Delete(&tasks, parsedId)
		// No tasks found is returned as an empty array
		// Let the client-side handle displaying of no tasks, if so
		zap.L().Info("Deleted task with id: " + id)
		// Send a HTTP 204 back since a succesful delete returns a null body
		return cxt.SendStatus(204)
	}
}
