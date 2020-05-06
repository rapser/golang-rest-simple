package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

type allTask []task

var tasks = allTask{
	{
		ID:      1,
		Name:    "task one",
		Content: "Some content",
	},
}

func getTasks(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request){

	var newTask task

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w,"Inserta un dato valido")
	}

	json.Unmarshal(reqBody, &newTask)
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

}

func getTask(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)

	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "invalid id")
		return
	}
	
	for _, task := range tasks {
		if task.ID == taskID {
			w.Header().Set("Content-Type","application/json");
			json.NewEncoder(w).Encode(task)

		}
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)

	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "invalid id")
		return
	}

	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i + 1:]...)
			fmt.Fprintf(w, "La tarea con id %v ha sido removido", taskID)
		}
	}

}

func updateTask(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	var updatedTask task

	if err != nil {
		fmt.Fprintf(w, "invalid id")
	}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "por favor ingrese data valida")
	}
	json.Unmarshal(reqBody, &updatedTask)

	for i, t := range tasks {
		if t.ID == taskID {
			tasks = append(tasks[:i], tasks[i + 1:]...)
			updatedTask.ID = taskID
			tasks = append(tasks, updatedTask)

			fmt.Fprintf(w, "La tarea con id %v ha sido actualizada", taskID)
		}
	}
}

func indexRoute(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "Bienvenido a mi API")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", router))
}
