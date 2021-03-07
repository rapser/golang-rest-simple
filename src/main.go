package main

import (
	"fmt"
	"log"
	"net/http"

	"firstgo_app/src/models"

	"github.com/gofiber/fiber/v2"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Bienvenido a mi API")
}

func setupRoutes(app *fiber.App) {

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/task", models.GetTasks)
	v1.Get("/task/:id", models.GetTask)
	v1.Post("/task", models.CreateTask)
	v1.Put("/task/:id", models.UpdateTask)
	v1.Delete("/task/:id", models.DeleteTask)
}

func main() {

	var port string = "3000"
	app := fiber.New()
	setupRoutes(app)

	// router := mux.NewRouter()
	// apiRouter := router.PathPrefix("/api/").Subrouter()

	// apiRouter.HandleFunc("/todos", api.GetTodos).Methods("GET")
	// apiRouter.HandleFunc("/todos", api.CreateTodo).Methods("POST")

	// apiRouter.HandleFunc("/todos/{id}", api.GetTodo).Methods("GET")
	// apiRouter.HandleFunc("/todos/{id}", api.DeleteTodo).Methods("DELETE")
	// apiRouter.HandleFunc("/todos/{id}", api.UpdateTodo).Methods("PUT")

	fmt.Printf("Server running at port %s", port)
	log.Fatal(app.Listen(":3000"))
}
