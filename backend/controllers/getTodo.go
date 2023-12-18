package controllers

import (
	"encoding/json"

	config "go-fiber-todo-backend/config"
	models "go-fiber-todo-backend/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func GetTodo(cxt *fiber.Ctx) error {
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
			return cxt.Status(500).JSON(fiber.Map{"err": err.Error()})
		}

		db.Find(&tasks, parsedId)
		if len(tasks) > 0 {
			// For now, marshal the return tasks struct into JSON and parse into a sring
			out, err := json.Marshal(&tasks)
			if err != nil {
				zap.L().Error(err.Error())
			}
			zap.L().Info("Found task with id: " + id)
			zap.L().Info(string(out))
			// No tasks found is returned as an empty array
			// Let the client-side handle displaying of no tasks, if so
			return cxt.JSON(fiber.Map{"msg": &tasks})
		} else {
			// No tasks found is returned as a 404 back to the client
			// Let the client-side handle displaying of no tasks, if so
			zap.L().Info("No task found with id: " + id)
			// Send a HTTP 404 back since nothing was found with this id
			return cxt.SendStatus(404)
		}
	}
}
