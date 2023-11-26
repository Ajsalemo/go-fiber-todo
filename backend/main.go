package main

import (
	config "go-fiber-todo-backend/config"
	"go-fiber-todo-backend/controllers"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func main() {
	app := fiber.New()
	api := app.Group("/api/todo")

	_, err := config.ConnectDB()
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	app.Get("/", controllers.Index)
	api.Get("/get", controllers.GetAllTodo)
	api.Get("/get/:id", controllers.GetTodo)
	api.Get("/create", controllers.CreateTodo)
	api.Get("/delete/:id", controllers.DeleteTodo)
	api.Get("/update/id", controllers.UpdateTodo)
	app.All("*", controllers.Index)

	app.Listen(":3000")

}
