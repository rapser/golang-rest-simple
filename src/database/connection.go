package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func GetConnection() *sql.DB {
// 	connStr := "postgres://postgres:root1983@localhost/todo_db?sslmode=disable"
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return db
// }

func GetConnection() *gorm.DB {

	dsn := "host=localhost user=postgres password=root1983 dbname=todo_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database connection successfully opened")

	return db
}
