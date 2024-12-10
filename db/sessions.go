package db

import "time"

type Session struct {
    ID        string
    UserID    int
    ExpiresAt time.Time
    CreatedAt time.Time
}