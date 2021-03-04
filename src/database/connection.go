package database

import (
	"database/sql"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	connStr := "postgres://postgres:root1983@localhost/todo_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetConnection2() *gorm.DB {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=todo_db sslmode=disable password=root1983")

	if err != nil {
		panic("failed to connect database")
	}

	return db

}
