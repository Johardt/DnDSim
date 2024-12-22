package db

import (
	"database/sql"
	"log"
	"sync"
)

var (
	DB   *sql.DB
	once sync.Once
)

func InitializeDB(dataSourceName string) {
	once.Do(func() {
		var err error
		DB, err = sql.Open("sqlite3", dataSourceName)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}

		setPragmas()

		err = createUsersTable()
		if err != nil {
			log.Fatalf("Failed to create users table: %v", err)
		}

		err = createSessionsTable()
		if err != nil {
			log.Fatalf("Failed to create sessions table: %v", err)
		}
	})
}

func setPragmas() {
	_, err := DB.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatalf("Failed to enable foreign keys: %v", err)
	}

	_, err = DB.Exec("PRAGMA synchronous = NORMAL;")
	if err != nil {
		log.Fatalf("Failed to set PRAGMA synchronous: %v", err)
	}

	_, err = DB.Exec("PRAGMA journal_mode = WAL;")
	if err != nil {
		log.Fatalf("Failed to set PRAGMA journal_mode: %v", err)
	}
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
