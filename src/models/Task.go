package models

import (
	"firstgo_app/src/database"

	"github.com/gofiber/fiber/v2"
)

// Task son las tareas presentes
type Task struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

// GetTasks obtiene todos las tareas
func GetTasks(c *fiber.Ctx) error {

	db := database.GetConnection()
	var tasks []Task
	db.Find(&tasks)
	return c.JSON(tasks)
}

// CreateTask crea una nueva tarea
func CreateTask(c *fiber.Ctx) error {

	db := database.GetConnection()
	task := new(Task)

	if err := c.BodyParser(task); err != nil {
		c.Status(503).Send([]byte(err.Error()))
	}

	db.Create(&task)
	return c.Status(201).JSON(task)
}

// GetTask obtiene una tarea segun su id
func GetTask(c *fiber.Ctx) error {

	db := database.GetConnection()
	id := c.Params("id")
	var task Task

	db.Find(&task, id)

	if task.ID == 0 {
		return c.Status(404).SendString("No encontrado")
	}
	return c.JSON(task)
}

// UpdateTask : sirve para actualizar una tarea
func UpdateTask(c *fiber.Ctx) error {

	db := database.GetConnection()

	var updateTask Task
	if err := c.BodyParser(&updateTask); err != nil {
		c.Status(503).Send([]byte(err.Error()))
	}

	var task Task
	id := c.Params("id")
	db.First(&task, id)

	db.Model(&task).Updates(Task{Name: updateTask.Name, Content: updateTask.Content})
	return c.JSON(task)

}

// DeleteTask : Sirve para eliminar una tarea
func DeleteTask(c *fiber.Ctx) error {

	db := database.GetConnection()
	id := c.Params("id")
	var task Task
	db.Delete(&task, id)
	return c.SendString("Task Successfully deleted")
}
