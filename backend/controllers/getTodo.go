package controllers

import (
	config "go-fiber-todo-backend/config"
	models "go-fiber-todo-backend/models"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetTodo(cxt *fiber.Ctx) error {
	id := cxt.Params("id")
	var tasks []models.Task

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	// If the request param is empty, return it to the client for potential error handling
	if id == "" {
		log.Default().Print("[WARN] Parameter must not be empty")
		return cxt.JSON(fiber.Map{"err": "Parameter must not be empty"})
	} else {
		// Parse the request param into an Int type
		parsedId, err := strconv.ParseInt(id, 10, 64)
		// If there is an issue with Int conversion with the request param, return it to the client for potential error handling
		if err != nil {
			return cxt.JSON(fiber.Map{"err": err})
		}

		db.Find(&tasks, parsedId)
		// No tasks found is returned as an empty array
		// Let the client-side handle displaying of no tasks, if so
		return cxt.JSON(fiber.Map{"msg": &tasks})
	}
}
