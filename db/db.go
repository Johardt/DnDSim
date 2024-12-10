package db

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB   *sql.DB
	once sync.Once
)

type User struct {
	ID        int
	Email     string
	Password  string
	CreatedAt string
}

func InitializeDB(dataSourceName string) {
	once.Do(func() {
		var err error
		DB, err = sql.Open("sqlite3", dataSourceName)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}

		// Create users table
		createTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`

		_, err = DB.Exec(createTableQuery)
		if err != nil {
			log.Fatalf("Failed to create users table: %v", err)
		}
	})
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

func CreateUser(email, hashedPassword string) error {
	query := `INSERT INTO users (email, password) VALUES (?, ?);`
	_, err := DB.Exec(query, email, hashedPassword)
	return err
}

func GetUserByEmail(email string) (*User, error) {
	query := `SELECT id, email, password, created_at FROM users WHERE email = ?;`
	row := DB.QueryRow(query, email)

	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
