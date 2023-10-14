package repo

import (
	"context"
	"database/sql"
	"time"
)

type Note struct {
	ID        string         `db:"id"`
	Title     sql.NullString `db:"title"`
	Text      sql.NullString `db:"text"`
	Author    sql.NullString `db:"author"`
	Email     sql.NullString `db:"email"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
}

type NoteRepository interface {
	Create(ctx context.Context, note *Note) (string, error)
	Get(ctx context.Context, id string) (*Note, error)
	GetList(ctx context.Context, query string, limit int64, offset int64) ([]*Note, error)
	Update(ctx context.Context, id string, note *Note) error
	Delete(ctx context.Context, id string) error
}
