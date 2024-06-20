package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func GetDatabase() *sql.DB {
	connStr := "postgresql://nuublx:psgrPass@localhost:5432/books_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
