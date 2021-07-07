package models

import "time"

type Book struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Abstract  string    `json:"abstract"`
	CreatedAt time.Time `json:"created_at"`
}
