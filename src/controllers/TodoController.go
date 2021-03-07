package controllers

import (
	"firstgo_app/src/database"
	"firstgo_app/src/models"

	"github.com/gofiber/fiber/v2"
)

// GetTodos obtiene todos las tareas
func GetTodos(c *fiber.Ctx) error {

	db := database.GetConnection()
	var todos []models.Todo
	db.Find(&todos)
	return c.JSON(todos)
}

// CreateTodo crea una nueva tarea
func CreateTodo(c *fiber.Ctx) error {

	db := database.GetConnection()
	todo := new(models.Todo)

	if err := c.BodyParser(todo); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"success": false,
			"message": "Json dont parser",
		})
	}

	db.Create(&todo)
	return c.Status(201).JSON(todo)
}

// GetTodo obtiene una tarea segun su id
func GetTodo(c *fiber.Ctx) error {

	db := database.GetConnection()
	id := c.Params("id")
	var todo models.Todo

	if result := db.Find(&todo, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Not found id",
		})
	}

	return c.JSON(todo)
}

// UpdateTodo : sirve para actualizar una tarea
func UpdateTodo(c *fiber.Ctx) error {

	db := database.GetConnection()

	var updateTodo models.Todo
	if err := c.BodyParser(&updateTodo); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"success": false,
			"message": "No se puede procesar",
		})
	}

	var todo models.Todo
	id := c.Params("id")

	if err := db.First(&todo, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Not found id",
		})
	}

	db.Model(&todo).Updates(models.Todo{Description: updateTodo.Description})
	return c.JSON(todo)

}

// DeleteTodo : Sirve para eliminar una tarea
func DeleteTodo(c *fiber.Ctx) error {

	db := database.GetConnection()
	id := c.Params("id")
	var todo models.Todo

	if err := db.Delete(&todo, id); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Not found id",
		})
	}
	return c.SendString("Todo Successfully deleted")
}
