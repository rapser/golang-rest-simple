package controllers

import (
	"firstgo_app/src/database"
	"firstgo_app/src/models"

	"github.com/gofiber/fiber/v2"
)

// GetTasks obtiene todos las tareas
func GetTasks(c *fiber.Ctx) error {

	db := database.GetConnection()
	var tasks []models.Task
	db.Find(&tasks)
	return c.JSON(tasks)
}

// CreateTask crea una nueva tarea
func CreateTask(c *fiber.Ctx) error {

	db := database.GetConnection()
	task := new(models.Task)

	if err := c.BodyParser(task); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"success": false,
			"message": "Json dont parser",
		})
	}

	db.Create(&task)
	return c.Status(201).JSON(task)
}

// GetTask obtiene una tarea segun su id
func GetTask(c *fiber.Ctx) error {

	db := database.GetConnection()
	id := c.Params("id")
	var task models.Task

	if result := db.Find(&task, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Not found id",
		})
	}

	return c.JSON(task)
}

// UpdateTask : sirve para actualizar una tarea
func UpdateTask(c *fiber.Ctx) error {

	db := database.GetConnection()

	var updateTask models.Task
	if err := c.BodyParser(&updateTask); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"success": false,
			"message": "No se puede procesar",
		})
	}

	var task models.Task
	id := c.Params("id")

	if err := db.First(&task, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Not found id",
		})
	}

	db.Model(&task).Updates(models.Task{Name: updateTask.Name, Content: updateTask.Content})
	return c.JSON(task)

}

// DeleteTask : Sirve para eliminar una tarea
func DeleteTask(c *fiber.Ctx) error {

	db := database.GetConnection()
	id := c.Params("id")
	var task models.Task

	if err := db.Delete(&task, id); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Not found id",
		})
	}
	return c.SendString("Task Successfully deleted")
}
