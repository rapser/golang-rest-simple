package main

import (
	"fmt"
	"log"
	"net/http"

	"firstgo_app/src/api"

	"github.com/gorilla/mux"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Bienvenido a mi API")
}

func main() {

	var port string = "3000"

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()

	apiRouter.HandleFunc("/todos", api.GetTodos).Methods("GET")
	apiRouter.HandleFunc("/todos", api.CreateTodo).Methods("POST")

	apiRouter.HandleFunc("/todos/{id}", api.GetTodo).Methods("GET")
	apiRouter.HandleFunc("/todos/{id}", api.DeleteTodo).Methods("DELETE")
	apiRouter.HandleFunc("/todos/{id}", api.UpdateTodo).Methods("PUT")

	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", indexRoute)
	// apiRouter.HandleFunc("/tasks", api.).Methods("GET")
	// apiRouter.HandleFunc("/tasks", Task.createTask).Methods("POST")
	// router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	// router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	// router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")

	fmt.Printf("Server running at port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
