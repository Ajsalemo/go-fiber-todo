package main

import (
	config "go-fiber-todo-backend/config"
	"go-fiber-todo-backend/controllers"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func main() {
	_, err := config.ConnectDB()
	// This helper function implements exponential retry backoffs for connection failure attempts to the database
	config.ConnectionRetry(err)

	app := fiber.New()
	api := app.Group("/api/todo")
	// Set up CORS to allow the frontends to call this API
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, DELETE, PUT, OPTIONS",
	}))

	app.Get("/", controllers.Index)
	api.Get("/get", controllers.GetAllTodo)
	api.Get("/get/:id", controllers.GetTodo)
	api.Post("/create", controllers.CreateTodo)
	api.Delete("/delete/:id", controllers.DeleteTodo)
	api.Put("/update/:id", controllers.UpdateTodo)
	app.All("*", controllers.Index)
	// Notify the application of the below signals to be handled on shutdown
	s := make(chan os.Signal, 1)
	signal.Notify(s,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	// Goroutine to clean up prior to shutting down
	go func() {
		sig := <-s
		switch sig {
		// Gorm should automatically handle connection closures. Close() is no longer available as of 1.20
		case os.Interrupt:
			zap.L().Warn("CTRL+C / os.Interrupt recieved, shutting down connections and the application..")
			app.Shutdown()
		case syscall.SIGTERM:
			zap.L().Warn("SIGTERM recieved.., shutting down connections and the application..")
			app.Shutdown()
		case syscall.SIGQUIT:
			zap.L().Warn("SIGQUIT recieved.., shutting down connections and the application..")
			app.Shutdown()
		case syscall.SIGINT:
			zap.L().Warn("SIGINT recieved.., shutting down connections and the application..")
			app.Shutdown()
		}
	}()

	app.Listen(":3000")
}
