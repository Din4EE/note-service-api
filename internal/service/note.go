package service

import "time"

type Note struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	Author    string    `json:"author"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
