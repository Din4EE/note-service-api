package model

import "time"

type Note struct {
	ID        uint64     `json:"id"`
	NoteInfo  *NoteInfo  `json:"info"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type NoteInfo struct {
	Title  *string `json:"title"`
	Text   *string `json:"text"`
	Author *string `json:"author"`
	Email  *string `json:"email"`
}
