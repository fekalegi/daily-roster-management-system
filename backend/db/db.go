package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
	return db
}

func Migrate(db *sql.DB) {
	content, err := os.ReadFile("db/schema.sql")
	if err != nil {
		log.Fatalf("Failed to read schema.sql: %v", err)
	}

	if _, err := db.Exec(string(content)); err != nil {
		log.Fatalf("Failed to apply migration: %v", err)
	}
}
