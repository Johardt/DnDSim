package db

import (
	"github.com/google/uuid"
	"time"
)

// TODO Implement maintenance routine to delete expired sessions

type Session struct {
	ID        string
	UserID    int
	ExpiresAt time.Time
	CreatedAt time.Time
}

func createSessionsTable() error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS sessions (
		id TEXT PRIMARY KEY,
		user_id INTEGER NOT NULL,
		created_at DATETIME NOT NULL,
		expires_at DATETIME NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);
	
	CREATE INDEX IF NOT EXISTS idx_sessions_user_id ON sessions(user_id);`

	_, err := DB.Exec(createTableQuery)
	return err
}

func CreateSession(userId int) (string, error) {
	sessionId := uuid.New().String()

	query := `INSERT INTO sessions (id, user_id, created_at, expires_at) VALUES (?, ?, ?, ?);`
	_, err := DB.Exec(query, sessionId, userId, time.Now(), time.Now().Add(time.Hour*24))
	if err != nil {
		return "", err
	}
	return sessionId, nil
}

func GetSessionByID(id string) (*Session, error) {
	query := `SELECT user_id, expires_at, created_at FROM sessions WHERE id = ?;`
	row := DB.QueryRow(query, id)

	var session Session
	err := row.Scan(&session.UserID, &session.ExpiresAt, &session.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func DeleteSession(id string) error {
	query := `DELETE FROM sessions WHERE id = ?;`
	_, err := DB.Exec(query, id)
	return err
}

func UpdateSessionExpiration(id string) error {
	query := `UPDATE sessions SET expires_at = ? WHERE id = ?;`
	_, err := DB.Exec(query, time.Now().Add(time.Hour*24), id)
	return err
}

func SessionExists(userId int) bool{
	query := `SELECT id FROM sessions WHERE user_id = ?;`
	row := DB.QueryRow(query, userId)

	var id string
	err := row.Scan(&id)
	return err == nil
}