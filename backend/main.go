package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"go-fiber-todo-backend/controllers"
)

func main() {
	app := fiber.New()
	api := app.Group("/api/todo")

	app.Get("/", controllers.Index)
	api.Get("/get/id", controllers.GetTodo)
	api.Get("/get/all", controllers.GetAllTodo)
	api.Get("/create", controllers.CreateTodo)
	api.Get("/delete/id", controllers.DeleteTodo)
	api.Get("/update/id", controllers.UpdateTodo)

	log.Fatal(app.Listen(":3000"))
}
