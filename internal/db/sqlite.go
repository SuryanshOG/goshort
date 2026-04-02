package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() *sql.DB {
	log.Println("🔥 Connecting to SQLite...")

	db, err := sql.Open("sqlite3", "goshort.db")
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	log.Println("✅ Connected to SQLite")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		original_url TEXT NOT NULL UNIQUE,
		short_code TEXT UNIQUE NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	return db
}
