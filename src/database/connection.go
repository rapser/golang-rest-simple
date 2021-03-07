package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetConnection : se conecta con una db postgres
func GetConnection() *gorm.DB {

	dsn := "host=localhost user=postgres password=root1983 dbname=todo_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database connection successfully opened")

	return db
}
