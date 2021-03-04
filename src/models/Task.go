package models

import (
	"encoding/json"
	"firstgo_app/src/database"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
)

// Task son las tareas presentes
type Task struct {
	gorm.Model

	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func getTasks(w http.ResponseWriter, r *http.Request) {

	db := database.GetConnection2()
	w.Header().Set("Content-Type", "application/json")

	var tasks []Task
	db.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {

	var newTask Task

	db := database.GetConnection2()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserta un dato valido")
	}

	json.Unmarshal(reqBody, &newTask)
	db.Create(&newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

}

// func getTask(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)

// 	taskID, err := strconv.Atoi(vars["id"])

// 	if err != nil {
// 		fmt.Fprintf(w, "invalid id")
// 		return
// 	}

// 	for _, task := range tasks {
// 		if task.ID == taskID {
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(task)

// 		}
// 	}
// }

// func deleteTask(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)

// 	taskID, err := strconv.Atoi(vars["id"])

// 	if err != nil {
// 		fmt.Fprintf(w, "invalid id")
// 		return
// 	}

// 	for i, task := range tasks {
// 		if task.ID == taskID {
// 			tasks = append(tasks[:i], tasks[i+1:]...)
// 			fmt.Fprintf(w, "La tarea con id %v ha sido removido", taskID)
// 		}
// 	}

// }

// func updateTask(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	taskID, err := strconv.Atoi(vars["id"])

// 	var updatedTask task

// 	if err != nil {
// 		fmt.Fprintf(w, "invalid id")
// 	}

// 	reqBody, err := ioutil.ReadAll(r.Body)

// 	if err != nil {
// 		fmt.Fprintf(w, "por favor ingrese data valida")
// 	}
// 	json.Unmarshal(reqBody, &updatedTask)

// 	for i, t := range tasks {
// 		if t.ID == taskID {
// 			tasks = append(tasks[:i], tasks[i+1:]...)
// 			updatedTask.ID = taskID
// 			tasks = append(tasks, updatedTask)

// 			fmt.Fprintf(w, "La tarea con id %v ha sido actualizada", taskID)
// 		}
// 	}
// }
