package controllers

import (
	config "go-fiber-todo-backend/config"
	models "go-fiber-todo-backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetAllTodo(cxt *fiber.Ctx) error {
	var tasks []models.Task

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	db.Find(&tasks)
	return cxt.JSON(fiber.Map{"data": &tasks})
}
