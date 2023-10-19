package model

import (
	"database/sql"
	"time"
)

type Note struct {
	ID        uint64         `db:"id"`
	Title     sql.NullString `db:"title"`
	Text      sql.NullString `db:"text"`
	Author    sql.NullString `db:"author"`
	Email     sql.NullString `db:"email"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
}
