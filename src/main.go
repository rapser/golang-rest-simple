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

	app.Get("/api/v1/task", models.GetTasks)
	app.Get("/api/v1/task/:id", models.GetTask)
	app.Post("/api/v1/task", models.CreateTask)
	app.Put("/api/v1/task/:id", models.UpdateTask)
	app.Delete("/api/v1/task/:id", models.DeleteTask)
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

	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", indexRoute)
	// apiRouter.HandleFunc("/tasks", api.).Methods("GET")
	// apiRouter.HandleFunc("/tasks", Task.createTask).Methods("POST")
	// router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	// router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	// router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")

	fmt.Printf("Server running at port %s", port)
	log.Fatal(app.Listen(":3000"))
}
