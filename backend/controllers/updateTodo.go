package controllers

import (
	"encoding/json"
	"fmt"
	config "go-fiber-todo-backend/config"
	models "go-fiber-todo-backend/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func UpdateTodo(cxt *fiber.Ctx) error {
	id := cxt.Params("id")
	var tasks models.Task

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
		// Try to find the task before attempting to update it
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
		// Update a task
		rows := db.Model(&tasks).Where("id = ?", parsedId).Updates(&tasks)
		fmt.Println(rows.RowsAffected)
		if rows.RowsAffected == 0 {
			zap.L().Info("Task with id: " + id + " was not found")
			// Send a HTTP 404 back since nothing was found with this id
			return cxt.SendStatus(404)
		} else if rows.RowsAffected == 1 {
			// Send a HTTP 204 back since a succesful update returns a null body
			zap.L().Info("Task updated with id: " + id)
			return cxt.SendStatus(204)
		} else if rows.Error != nil {
			zap.L().Error(rows.Error.Error())
			return cxt.Status(500).JSON(fiber.Map{"err": rows.Error.Error()})
		} else {
			zap.L().Error("An unknown error has occurred")
			return cxt.Status(500).JSON(fiber.Map{"err": "An unknown error has occurred"})
		}
	}
}
