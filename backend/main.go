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
	api.Get("/delete", controllers.DeleteTodo)
	api.Get("/update", controllers.UpdateTodo)

	log.Fatal(app.Listen(":3000"))
}
