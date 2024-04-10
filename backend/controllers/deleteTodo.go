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
	// This helper function implements exponential retry backoffs for connection failure attempts to the database
	config.ConnectionRetry(err)
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
			return cxt.Status(500).JSON(fiber.Map{"err": err.Error()})
		}
		// Try to find the task before attempting to delete it
		db.First(&tasks, parsedId)
		if len(tasks) > 0 {
			zap.L().Info("Found task with id: " + id + " to be deleted")
			err := db.Delete(&tasks)
			// Log out the error and return a 500 if the task can't be deleted
			if err.Error != nil {
				zap.L().Error(db.Error.Error())
				return cxt.SendStatus(500)
			}
			zap.L().Info("Deleted task with id: " + id)
			// Send a HTTP 204 back since a succesful delete returns a null body
			return cxt.SendStatus(204)
		} else {
			// No tasks found is returned as a 404 back to the client
			// Let the client-side handle displaying of no tasks, if so
			zap.L().Info("No task found with id: " + id)
			// Send a HTTP 404 back since nothing was found with this id
			return cxt.SendStatus(404)
		}
	}
}
