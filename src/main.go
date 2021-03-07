package main

import (
	"firstgo_app/src/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Bienvenido a mi API")
}

func setupRoutes(app *fiber.App) {

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/task", controllers.GetTasks)
	v1.Get("/task/:id", controllers.GetTask)
	v1.Post("/task", controllers.CreateTask)
	v1.Put("/task/:id", controllers.UpdateTask)
	v1.Delete("/task/:id", controllers.DeleteTask)

	v1.Get("/todo", controllers.GetTodos)
	v1.Post("/todo", controllers.CreateTodo)
	v1.Get("/todo/:id", controllers.GetTodo)
	v1.Delete("/todo/:id", controllers.DeleteTodo)
	v1.Put("/todo/:id", controllers.UpdateTodo)
}

func main() {

	var port string = "3000"
	app := fiber.New()
	setupRoutes(app)

	fmt.Printf("Server running at port %s", port)
	log.Fatal(app.Listen(":3000"))
}
