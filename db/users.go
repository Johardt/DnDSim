package db

import (
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	CreatedAt string
}

func createUsersTable() error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(createTableQuery)
	return err
}

func CreateUser(username, email, hashedPassword string) error {
	query := `INSERT INTO users (email, password) VALUES (?, ?);`
	_, err := DB.Exec(query, email, hashedPassword)
	return err
}

func GetUserByName(username string) (*User, error) {
	query := `SELECT id, username, email, password, created_at FROM users WHERE username = ?;`
	row := DB.QueryRow(query, username)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
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